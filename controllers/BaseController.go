package controllers

import (
	"github.com/astaxie/beego"
)

type BaseController struct {
	beego.Controller
}

func (c *BaseController) Prepare() {
	c.Layout = "layout/main.tpl"

	session := c.StartSession()
	userId := session.Get("userId")

	c.Data["isLoggedIn"] = userId != nil

	if userId == nil && c.routeNeedsAuth() {
		c.Redirect("/user/login", 302)
		return
	}
}

func (c *BaseController) routeNeedsAuth() bool {
	if c.Ctx.Request.URL.Path == "/user/login" ||
		c.Ctx.Request.URL.Path == "/user/register" {
		return false
	}

	return true
}
