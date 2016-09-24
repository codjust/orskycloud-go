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
	//flash := beego.NewFlash()
	//username, password := c.Input().Get("username"), c.Input().Get("password")
	username, password := c.GetString("username"), c.GetString("password")

	c.Data["json"] = "error"
	c.ServeJSONP()
	beego.Debug("username:", username, password)

	//c.Redirect("/", 302)

}
