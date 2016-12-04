package controllers

import (
	"github.com/astaxie/beego"
	//"orskycloud-go/cache_module"
	//"orskycloud-go/logicfunc"
	"orskycloud-go/models"
	"os"
	"strconv"
)

type SensorController struct {
	beego.Controller
}

// "createTime": "2016-9-12 00:00:00",
//             "designation": "舒张压",
//             "name": "diastolic pressure",
//             "unit": "mmHg"

func (this *SensorController) MySensor() {
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

func (this *SensorController) NewSensor() {
	//this.Data["Active_Sensor"] = "active"
	username := this.GetSession("username")
	this.Layout = "layout/layout.tpl"
	this.TplName = "newsensor.tpl"
	this.LayoutSections = make(map[string]string)
	//this.LayoutSections["Scripts"] = "scripts/my_sensor_scripts.tpl"
	this.Data["User"] = username

}
