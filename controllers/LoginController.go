package controllers

import (
	"github.com/astaxie/beego"
	"orskycloud/models"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Login() {
	c.TplName = "login.html"
}

func (c *LoginController) Register() {
	c.TplName = "register.html"
}

func (c *LoginController) RegisterInfo() {
	username, password := c.GetString("username"), c.GetString("password")
	info, err := models.AddUser(username, password)
	if err != nil {
		beego.Error("regist failed")
	}
	result := struct {
		Val string
	}{info}
	c.Data["json"] = &result
	c.ServeJSON()
	beego.Debug("username:", username, password, info)
}
