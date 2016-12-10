package controllers

import (
	"github.com/astaxie/beego"
	//"orskycloud-go/cache_module"
	//"orskycloud-go/logicfunc"
	//	"orskycloud-go/models"
	//	"os"
	//"strconv"
	//"time"
)

type HistoryController struct {
	beego.Controller
}

func (this *HistoryController) HistoryPage() {

	this.TplName = "historydata.tpl"
	this.Data["Active_History"] = "active"
	this.Layout = "layout/layout.tpl"
	this.LayoutSections = make(map[string]string)
	//this.LayoutSections["Scripts"] = "scripts/my_sensor_scripts.tpl"
	//this.Data["User"] = username
}
