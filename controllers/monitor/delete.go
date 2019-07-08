package monitor

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/domain"
	"github.com/astaxie/beego"
	"strconv"
)

type DeleteMonitorController struct {
	controllers.BaseController
}

func (c *DeleteMonitorController) Get() {
	c.TplName = "user/monitor/delete.tpl"

	monId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	monitor, _ := domain.GetMonitorService().Find(monId)

	if !c.IsAuthorized(monitor.UserId) {
		return
	}

	c.Data["monitor"] = monitor
}

func (c *DeleteMonitorController) Post() {
	monId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	monitor, _ := domain.GetMonitorService().Find(monId)

	if !c.IsAuthorized(monitor.UserId) {
		return
	}

	go c.deleteMonitor(monId)

	flash := beego.NewFlash()
	flash.Error("Monitor schedules for deletion")
	flash.Store(&c.Controller)

	c.Redirect("/user", 302)
}

func (c *DeleteMonitorController) deleteMonitor(monId int) {
	domain.GetWorkerService().Stop(monId)
	domain.GetMonitorService().Delete(monId)
}

