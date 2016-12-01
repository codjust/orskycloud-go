package routers

import (
	"orskycloud-go/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//beego.Router("/", &controllers.VerifyController{})
	beego.Router("/", &controllers.LoginController{}, "GET:Login")
	beego.Router("/register", &controllers.LoginController{}, "GET:Register")
	beego.Router("/register/handler", &controllers.LoginController{}, "POST:RegisterInfo")
	beego.Router("/login/handler", &controllers.LoginController{}, "POST:LoginCheck")
	beego.Router("/homepage", &controllers.HomePageController{}, "GET:HomePage")
	beego.Router("/mydevice/?:page", &controllers.HomePageController{}, "GET:MyDevice")

	beego.Router("/test", &controllers.HomePageController{}, "GET:MyCache")
}
