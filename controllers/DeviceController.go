package controllers

import (
	"github.com/astaxie/beego"
	"orskycloud-go/models"
	"os"
	"strconv"
)

type DeviceController struct {
	beego.Controller
}

func (this *DeviceController) MyDevice() {
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
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

func (this *DeviceController) NewDevice() {
	//	this.Data["Page"] = page
	//	this.Data["Active_Dev"] = "active"
	username := this.GetSession("username").(string)
	this.Layout = "layout/layout.tpl"
	this.TplName = "newdevice.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/newdevice_scripts.tpl"
	this.Data["User"] = username
}

func (this *DeviceController) CreateDevice() {
	beego.Debug("XXXXXXX")
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	var newDevice models.Device
	newDevice.DevName = this.GetString("devicename")
	newDevice.Description = this.GetString("description")

	beego.Debug("XXXXXXX")
	res := models.CreateNewDevice(username, password, newDevice)
	result := struct {
		Val string
	}{res}
	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *DeviceController) DeleteDevice() {
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	did := this.GetString("did")
	res := models.DeleteDeviceOp(username, password, did)
	result := struct {
		Val string
	}{res}
	this.Data["json"] = &result
	this.ServeJSON()
}

func (this *DeviceController) EditDevice() {
	did := this.Ctx.Input.Param(":did")
	beego.Debug("did:", did)

	username, password := this.GetSession("username").(string), this.GetSession("password").(string)

	ret_data := models.ReturnByIdDeviceInfo(username, password, did)

	ret_data.ID = did
	beego.Debug("ret_data:", ret_data)
	this.Data["DeviceName"] = ret_data.DevName
	this.Data["Description"] = ret_data.Description
	this.Data["Did"] = ret_data.ID
	this.Layout = "layout/layout.tpl"
	this.TplName = "editdevice.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/editdevice_scripts.tpl"
	this.Data["User"] = username

}

func (this *DeviceController) EditDeviceModify() {
	var dev models.Device
	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	dev.DevName, dev.Description = this.GetString("devicename"), this.GetString("description")
	dev.ID = this.GetString("did")
	res := models.UpdateDeviceInfo(username, password, dev)
	result := struct {
		Val string
	}{res}
	this.Data["json"] = &result
	this.ServeJSON()
}
