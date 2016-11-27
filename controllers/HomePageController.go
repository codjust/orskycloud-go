package controllers

import (
	"github.com/astaxie/beego"
	"orskycloud-go/logicfunc"
)

type HomePageController struct {
	beego.Controller
}

func (this *HomePageController) HomePage() {
	//这里要判断一下是否登录isLogin

	this.SetSession("username", "john")
	this.SetSession("password", "123456")

	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	last_logic_time := logicfunc.GetHomePage(username, password)
	this.Data["Last_login_time"] = last_logic_time
	this.Layout = "layout/layout.tpl"
	this.TplName = "homepage.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/home_scripts.tpl"
}
