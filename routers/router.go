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

	beego.Router("/mydevice/?:page", &controllers.DeviceController{}, "GET:MyDevice")
	beego.Router("/mydevice/newdevice", &controllers.DeviceController{}, "GET:NewDevice")
	beego.Router("/mydevice/create", &controllers.DeviceController{}, "POST:CreateDevice")
	beego.Router("/mydevice/delete", &controllers.DeviceController{}, "POST:DeleteDevice")
	beego.Router("/mydevice/edit/?:did", &controllers.DeviceController{}, "GET:EditDevice")
	beego.Router("/mydevice/edit/modify", &controllers.DeviceController{}, "POST:EditDeviceModify")

	beego.Router("/mysensor/?:page", &controllers.SensorController{}, "GET:MySensor")
	beego.Router("/mysensor/newdevice", &controllers.SensorController{}, "GET:NewSensor")
	beego.Router("/mysensor/new/create", &controllers.SensorController{}, "POST:CreateSensor")
	beego.Router("/mysensor/delete", &controllers.SensorController{}, "POST:DeleteSensor")

	beego.Router("/myprofile", &controllers.ProfileController{}, "GET:MyProfile")
	beego.Router("/myprofile/update", &controllers.ProfileController{}, "POST:Update")
	//beego.Router("/myprofile/check", &controllers.ProfileController{}, "POST:Check")
	beego.Router("/updatepwd", &controllers.ProfileController{}, "GET:UpdatePwd")
	beego.Router("/updatepwd/modify", &controllers.ProfileController{}, "POST:UpdatePwdModify")
	beego.Router("/test", &controllers.HomePageController{}, "GET:MyCache")
}
