package User

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/domain"
	"github.com/astaxie/beego"
)

type RegisterController struct {
	controllers.BaseController
}

func (c *RegisterController) Get() {
	c.TplName = "user/register.tpl"
}

func (c *RegisterController) Post() {
	c.TplName = "user/register.tpl"

	email := c.GetString("email")

	if domain.GetUserService().EmailExists(email) {
		flash := beego.NewFlash()
		flash.Error("Email already exists")
		flash.Store(&c.Controller)

		c.Layout = "layout/main.tpl"
		c.TplName = "user/register.tpl"
		return
	}

	password := c.GetString("password")
	passwordRepeat := c.GetString("password-repeat")

	if password != passwordRepeat {
		flash := beego.NewFlash()
		flash.Error("Password and confirmation are not equal")
		flash.Store(&c.Controller)

		c.Layout = "layout/main.tpl"
		c.TplName = "user/register.tpl"
		return
	}

	user := domain.GetUserService().Create(email, password)
	domain.GetUserService().Add(user)

	c.SetSession("userId", user.Id)

	c.Redirect("/user", 302)
}
