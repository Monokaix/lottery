package main

import (
	"myproject/lottery/bootstrap"
	"myproject/lottery/web/routes"
	"fmt"
	"myproject/lottery/web/middleware/identity"
)

var port = 8080

 func newApp() *bootstrap.Bootstrapper {
	 app := bootstrap.New("Go抽奖系统",
		 "Monokaix")
	 app.Bootstrap()
	 app.Configure(identity.Configure, routes.Configure)
	 return app
 }

 func main(){
 	app := newApp()
 	app.Listen(fmt.Sprintf("%d", port))
 }