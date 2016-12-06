package controllers

import (
	"github.com/astaxie/beego"
	//"orskycloud-go/cache_module"
	//"orskycloud-go/logicfunc"
	"orskycloud-go/models"
	"os"
	"strconv"
	"time"
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
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)

	d_list := models.ReturnDevList(username, password)

	this.Data["DList"] = d_list
	this.Layout = "layout/layout.tpl"
	this.TplName = "newsensor.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/newsensor_scripts.tpl"
	this.Data["User"] = username

}

func (this *SensorController) CreateSensor() {
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	localtime := time.Now().Format("2006-01-02 15:04:05")
	var new_sensor models.Sensor
	new_sensor.Name = this.GetString("name")
	new_sensor.Designation = this.GetString("designation")
	new_sensor.Unit = this.GetString("unit")
	new_sensor.Did = this.GetString("did")
	new_sensor.CreateTime = localtime
	res := models.CreateNewSensor(username, password, new_sensor)
	result := struct {
		Val string
	}{res}
	this.Data["json"] = &result
	this.ServeJSON()
}
