package controllers

import "UptimeMonitor/domain"

type DebugController struct {
	BaseController
}

func (c *DebugController) Get() {
	c.TplName = "debug.tpl"
	userId := c.GetSession("userId").(int)

	c.Data["allmons"] = domain.GetMonitorService().FindAll()
	c.Data["usermons"] = domain.GetMonitorService().FindUserMonitors(userId)
	c.Data["count"] = domain.GetUserMonitorsRepository(userId).Count()
}