package controllers

import (
	"github.com/astaxie/beego"
	"orskycloud-go/logicfunc"
	"orskycloud-go/models"
	"strconv"
)

type HomePageController struct {
	beego.Controller
}

func (this *HomePageController) HomePage() {
	//这里要判断一下是否登录isLogin

	this.SetSession("username", "john")
	this.SetSession("password", "123456")

	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	beego.Debug(username, password)
	last_logic_time := logicfunc.GetHomePage(username, password)
	beego.Debug("time:", last_logic_time)
	this.Data["Last_login_time"] = last_logic_time
	this.Data["User"] = username
	this.Layout = "layout/layout.tpl"
	this.TplName = "homepage.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/home_scripts.tpl"
}

func (this *HomePageController) MyDevice() {
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	// devices := models.ReturnAllDevices(username, password)
	// beego.Debug(devices)

	p, _ := strconv.Atoi(this.Ctx.Input.Query("p"))
	if p == 0 {
		p = 1
	}
	size, _ := beego.AppConfig.Int("page.size")
	page := models.PageUser(p, size, username, password)
	// this.Data["Devices"] = devices
	this.Data["Page"] = page
	this.Layout = "layout/layout.tpl"
	this.TplName = "my_device.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/my_device_scripts.tpl"
	this.Data["User"] = username
}
