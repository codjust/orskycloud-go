package controllers

import (
	"github.com/astaxie/beego"
	"orskycloud-go/models"
)

type ProfileController struct {
	beego.Controller
}

func (this *ProfileController) MyProfile() {

	username, password := this.GetSession("username").(string), this.GetSession("password").(string)
	profile := models.ReturnProfileInfo(username, password)

	this.Data["Profile"] = profile
	this.Data["Active_Profile"] = "active"
	this.Layout = "layout/layout.tpl"
	this.TplName = "my_profile.tpl"
	this.LayoutSections = make(map[string]string)
	this.LayoutSections["Scripts"] = "scripts/profile_script.tpl"
	this.Data["User"] = username
}

func (this *ProfileController) Update() {
	username, phone, email := this.GetString("username"), this.GetString("phone"), this.GetString("email")
	user, pwd := this.GetSession("username").(string), this.GetSession("password").(string)
	profile := models.Profile{UserName: username, Phone: phone, EMail: email}

	res := models.UpdataProfileInfo(user, pwd, profile)
	if res == "success" {
		//更新session
		this.SetSession("username", username)
	}
	result := struct {
		Val string
	}{res}
	this.Data["json"] = &result
	this.ServeJSON()
}
