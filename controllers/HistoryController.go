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

func (this *HistoryController) GetSensorList() {
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	Did := this.GetString("did")
	beego.Debug("did:", Did)
	ret_data := models.GetSenSor(username, password, Did)
	beego.Debug("ret_data", ret_data)
	this.Data["json"] = &ret_data
	this.ServeJSON()
}

func (this *HistoryController) GetHistoryData() {
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	Did := this.GetString("did")
	Name := this.GetString("name")
	Start := this.GetString("start")
	End := this.GetString("end")

	models.GetHistory(username, password, Did, Name, Start, End)

}
