package controllers

import (
	"github.com/astaxie/beego"
	"orskycloud-go/cache_module"
	"orskycloud-go/logicfunc"
	"orskycloud-go/models"
	"os"
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
	this.TplName = "homepage.html"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/home_scripts.tpl"
}

func (this *HomePageController) MyDevice() {
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	// devices := models.ReturnAllDevices(username, password)
	// beego.Debug(devices)

	beego.Debug("page:", this.Ctx.Input.Param(":page"))
	var pageNum int
	var err error
	if this.Ctx.Input.Param(":page") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(this.Ctx.Input.Param(":page"))
		if err != nil {
			beego.Debug("error:", err)
			os.Exit(1)
		}
	}
	page := models.PageDevice(pageNum, username, password)
	// this.Data["Devices"] = devices
	this.Data["Page"] = page
	this.Data["Active_Dev"] = "active"
	this.Layout = "layout/layout.tpl"
	this.TplName = "my_device.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/my_device_scripts.tpl"
	this.Data["User"] = username

}

func (this *HomePageController) MyCache() {
	cache_module.PutData()
	beego.Debug("data:", cache_module.Get())
}

func (this *HomePageController) MySensor() {
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	//this.Data["Page"] = page
	var pageNum int
	var err error
	if this.Ctx.Input.Param(":page") == "" {
		pageNum = 1
	} else {
		pageNum, err = strconv.Atoi(this.Ctx.Input.Param(":page"))
		if err != nil {
			beego.Debug("error:", err)
			os.Exit(1)
		}
	}
	sensors := models.PageSensor(pageNum, username, password)
	this.Data["Page"] = sensors
	this.Data["Active_Sensor"] = "active"
	this.Layout = "layout/layout.tpl"
	this.TplName = "my_sensor.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/my_sensor_scripts.tpl"
	this.Data["User"] = username
}
