package comm

import (
	"net/http"
	"net"
	"myproject/lottery/models"
	"net/url"
	"strconv"
	"fmt"
	"myproject/lottery/conf"
	"crypto/md5"
	"github.com/lunny/log"
)

func ClientIP(request *http.Request) string{
	host, _, _ := net.SplitHostPort(request.RemoteAddr)
	return host
}

func Redirect(writer http.ResponseWriter, url string) {
	writer.Header().Add("Location",url)
	writer.WriteHeader(http.StatusFound)
}

//获取cookie
func GetLoginUser(request *http.Request) *models.ObjLoginuser  {
	c, err := request.Cookie("lottery_loginuser")
	if err != nil{
		return nil
	}
	params, err := url.ParseQuery(c.Value)
	if err != nil{
		return nil
	}
	uid, err := strconv.Atoi(params.Get("uid"))
	if err != nil || uid < 1 {
		return nil
	}
	now, err := strconv.Atoi(params.Get("now"))
	if err != nil || NowUnix() -now > 86400*30{
		return nil
	}

	loginuser := &models.ObjLoginuser{}
	loginuser.Uid = uid
	loginuser.Username = params.Get("username")
	loginuser.Now = now
	loginuser.Ip = ClientIP(request)
	loginuser.Sign = params.Get("sign")
	//判断签名是否正确
	sign := createLoginuserSign(loginuser)
	if sign != loginuser.Sign{
		log.Println("func_web GetLoginuser createLoginuserSign not signed",
			sign,loginuser.Sign)
		return nil
	}
	return loginuser
}

//设置cookie
func SetLoginuser(writer http.ResponseWriter, loginuser *models.ObjLoginuser) {
	if loginuser == nil || loginuser.Uid < 1 {
		//清除cookie
		c := &http.Cookie{
			Name :"lottery_loginuser",
			Value:"",
			Path:"/",
			MaxAge:-1,
		}
		http.SetCookie(writer,c)
		return
	}
	if loginuser.Sign == "" {
		loginuser.Sign = createLoginuserSign(loginuser)
	}
	//构造参数
	params := url.Values{}
	params.Add("uid",strconv.Itoa(loginuser.Uid))
	params.Add("username",loginuser.Username)
	params.Add("now",strconv.Itoa(loginuser.Now))
	params.Add("ip",loginuser.Ip)
	params.Add("sign",loginuser.Sign)
	//设置有效的cookie
	c := &http.Cookie{
		Name:"lottery_loginuser",
		Value:params.Encode(),
		Path:"/",
	}
	http.SetCookie(writer,c)

}
//创建签名
func createLoginuserSign(loginuser *models.ObjLoginuser) string {
	str := fmt.Sprintf("uid=%d&username=%ssecret=%s&now=%d",loginuser.Uid,
		loginuser.Username,conf.CookieSecret,loginuser.Sign)
	sign := fmt.Sprintf("%x",md5.Sum([]byte(str)))
	return sign
}