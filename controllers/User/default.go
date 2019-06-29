package User

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/dataAccess/repositories"
	"UptimeMonitor/domain/types"
)

type UserController struct {
	controllers.BaseController
}

func (c *UserController) Get() {
	c.Layout = "layout/main.tpl"
	c.TplName = "user/user_dashboard.tpl"

	c.Data["user"] = getUser()
	c.Data["monitors"] = getMonitors()
}

func getMonitors() interface{} {
	// @todo Fake
	return repositories.NewUserMonitorsFake().Get()
}

func getUser() types.User {
	// @todo Fake
	user := types.User{
		Id:    1,
		Email: "farzan.b@gmail.com",
		Name:  "Farzan",
	}

	return user
}