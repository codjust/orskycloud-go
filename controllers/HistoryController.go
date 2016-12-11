package controllers

import (
	"github.com/astaxie/beego"
	//"orskycloud-go/cache_module"
	//"orskycloud-go/logicfunc"
	"orskycloud-go/models"
	//	"os"
	//"strconv"
	//"time"
)

type HistoryController struct {
	beego.Controller
}

func (this *HistoryController) HistoryPage() {

	username, password := this.GetSession("username").(string), this.GetSession("password").(string)

	exp_data := models.GetDevSenList(username, password)
	this.Data["Data"] = exp_data
	this.TplName = "historydata.tpl"
	this.Data["Active_History"] = "active"
	this.Layout = "layout/layout.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/history_scripts.tpl"
	this.Data["User"] = username
}
