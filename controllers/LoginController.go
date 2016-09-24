package controllers

import (
	"github.com/astaxie/beego"
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
	result := struct {
		Val string
	}{username}
	c.Data["json"] = &result
	c.ServeJSON()
	beego.Debug("username:", username, password)

}
