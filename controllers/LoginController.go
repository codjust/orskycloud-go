package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/fzzy/radix/redis"
	"time"
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

	// client, err := redis.DialTimeout("tcp", "127.0.0.1:6379", time.Duration(10)*time.Second)
	// errHndlr(err)

	// r := client.Cmd("select", 0)
	// errHndlr(r.Err)

	// client.Cmd("set", "test", "sds")
	// if username != nil || password |= nil {
	// 	beego.Debug("username:", username, password, info)
	// }
	// result := struct {
	// 	Val string
	// }{info}
	// c.Data["json"] = &result
	// c.ServeJSON()
	beego.Debug("username:", username, password)
}
