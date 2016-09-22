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
