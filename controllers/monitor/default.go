package monitor

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/domain"
)

type DefaultController struct {
	controllers.BaseController
}

func (c *DefaultController) Get() {
	c.TplName = "user/user_dashboard.tpl"

	c.Data["user"] = c.getUser()
	c.Data["monitors"] = c.getMonitors()
}

func (c *DefaultController) getMonitors() interface{} {
	return domain.GetMonitorService().FindUserMonitors(c.getUserId())
}

func (c *DefaultController) getUser() domain.User {
	userId := c.GetSession("userId").(int)
	user, _ := domain.GetUserService().Get(userId)
	return user
}

func (c *DefaultController) getUserId() int {
	return c.GetSession("userId").(int)
}
