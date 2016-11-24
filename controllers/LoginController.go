package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"orskycloud-go/models"
)

type LoginController struct {
	beego.Controller
}

func errHndlr(err error) {
	if err != nil {
		fmt.Println("error:", err)
	}
}

func (c *LoginController) Login() {
	c.TplName = "login.html"
}

func (c *LoginController) Register() {
	c.TplName = "register.html"
}

func (c *LoginController) RegisterInfo() {
	username, password := c.GetString("username"), c.GetString("password")

	res := models.HandleRegist(username, password)

	result := struct {
		Val string
	}{res}
	c.Data["json"] = &result
	c.ServeJSON()
	beego.Debug("username:", res, username, password)
}
