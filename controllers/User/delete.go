package User

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/domain"
	"strconv"
)

type DeleteMonitorController struct {
	controllers.BaseController
}

func (c *DeleteMonitorController) Get() {
	c.Layout = "layout/main.tpl"
	c.TplName = "user/monitor/delete.tpl"

	monId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	monitor, _ := domain.GetMonitorService().Find(monId)

	c.Data["monitor"] = monitor
}

func (c *DeleteMonitorController) Post() {
	monId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

	domain.GetWorkerService().Stop(monId)
	domain.GetMonitorService().Delete(monId)

	c.Redirect("/user", 302)
}
