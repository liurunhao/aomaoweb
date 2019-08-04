package routers

import (
	"shop/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.MainController{},"get:ShowLogin;post:HandleLogin")
	beego.Router("/register", &controllers.MainController{},"get:ShowRegister;post:HandleRegister")
	beego.Router("/logout", &controllers.MainController{},"post:Logout")
	beego.Router("/cart", &controllers.MainController{}, "GET:Cart")
	beego.Router("/cartAddress", &controllers.MainController{}, "get:CartAddress")
	beego.Router("/orders", &controllers.MainController{}, "GET:Orders")

}
