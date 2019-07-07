package User

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/domain"
	"github.com/astaxie/beego"
	"golang.org/x/crypto/bcrypt"
)

type LoginController struct {
	controllers.BaseController
}

func (c *LoginController) Get() {
	c.TplName = "user/login.tpl"
}

func (c *LoginController) Post() {
	email := c.GetString("email")
	password := c.GetString("password")

	user, err := domain.GetUserService().GetByEmail(email)
	if err != nil {
		flash := beego.NewFlash()
		flash.Error("Email not found")
		flash.Store(&c.Controller)

		c.TplName = "user/login.tpl"
		return;
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
		flash := beego.NewFlash()
		flash.Error("Password is incorrect")
		flash.Store(&c.Controller)

		c.TplName = "user/login.tpl"
		return;
	}

	c.SetSession("userId", user.Id)
	c.Redirect("/user", 302)
}
