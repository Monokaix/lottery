package controllers

import (
	"github.com/kataras/iris"
	"myproject/lottery/services"
	"myproject/lottery/models"
)

type IndexController struct {
	Ctx iris.Context
	ServiceUser services.UserService
	ServiceBlackip services.BlackipService
	ServiceCode services.CodeService
	ServiceGift services.GiftService
	ServiceResult services.ResultService
	ServiceUserday services.UserdayService
}

//http:localhost:8080
func (c *IndexController) Get() string {
	c.Ctx.Header("Content-Type","text/html")
	return "Welcome to go 抽奖系统, <a href='/public/index.html'>开始抽奖</a>"
}

//http:localhost:8080/gifts
func (c *IndexController) GetGifts() map[string]interface{} {
	rs := make(map[string]interface{},0)
	rs["code"] = 0
	rs["msg"] = ""
	datalist := c.ServiceGift.GetAll()
	list := make([]models.LtGift,0)
	for _, data := range datalist{
		if data.Sysstatus == 0 {
			list = append(list,data)
		}
	}

	return rs
}

func (c *IndexController) GetNewprize() map[string]interface{} {
	rs := make(map[string]interface{},0)
	rs["code"] = 0
	rs["msg"] = ""
	// TODO:

	return rs
}