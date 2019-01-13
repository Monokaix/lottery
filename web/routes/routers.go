package routes

import (
	"myproject/lottery/bootstrap"
	"myproject/lottery/services"
	"github.com/kataras/iris/mvc"
	"myproject/lottery/web/controllers"
	"myproject/lottery/web/middleware"
)

func Configure(b *bootstrap.Bootstrapper)  {
	userService := services.NewUserService()
	codeService := services.NewCodeService()
	giftService := services.NewGiftService()
	blackipService := services.NewBlackipService()
	resultService := services.NewResultService()
	userdayService := services.NewUserdayService()

	index := mvc.New(b.Party("/"))
	index.Register(
		userService,
		codeService,
		giftService,
		blackipService,
		resultService,
		userdayService)

	index.Handle(new(controllers.AdminController))

	admin := mvc.New(b.Party("/admin"))
	//验证admin username password
	admin.Router.Use(middleware.BasicAuth)
	admin.Register(
		userService,
		codeService,
		giftService,
		blackipService,
		resultService,
		userdayService)

	admin.Handle(new(controllers.AdminController))
}
