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
}
