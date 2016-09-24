package routers

import (
	"orskycloud/controllers"

	"github.com/astaxie/beego"
)

func init() {

	//beego.Router("/", &controllers.VerifyController{})
	beego.Router("/", &controllers.LoginController{}, "GET:Login")
	beego.Router("/register", &controllers.LoginController{}, "GET:Register")
	beego.Router("/register/handler", &controllers.LoginController{}, "POST:RegisterInfo")
}
