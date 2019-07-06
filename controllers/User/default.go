package User

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/domain"
)

type DefaultController struct {
	controllers.BaseController
}

func (c *DefaultController) Get() {
	c.Layout = "layout/main.tpl"
	c.TplName = "user/user_dashboard.tpl"

	c.Data["user"] = getUser()
	c.Data["monitors"] = getMonitors()
}

func getMonitors() interface{} {
	return domain.GetMonitorService().FindAll()
}

func getUserId() int {
	// @todo Fake
	return 0;
}

func getUser() domain.User {
	// @todo Fake
	user := domain.User{
		Id:    1,
		Email: "farzan.b@gmail.com",
		Name:  "Farzan",
	}

	return user
}