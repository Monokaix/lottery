package controllers

import (
	"github.com/kataras/iris"
	"myproject/lottery/services"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	Ctx iris.Context
	ServiceUser services.UserService
	ServiceBlackip services.BlackipService
	ServiceCode services.CodeService
	ServiceGift services.GiftService
	ServiceResult services.ResultService
	ServiceUserday services.UserdayService
}

func (c *AdminController) Get() mvc.Result{
	return mvc.View{
		Name:"admin/index.html",
		Data:iris.Map{
			"Title":"管理后台",
			"Channel":"",
		},
		Layout:"admin/layout.html",
	}
}