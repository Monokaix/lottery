package routes

import (
	"myproject/lottery/bootstrap"
	"myproject/lottery/services"
	"github.com/kataras/iris/mvc"
	"myproject/lottery/web/controllers"
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

	index.Handle(new(controllers.IndexController))
}
