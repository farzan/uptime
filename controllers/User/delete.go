package User

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
	c.Layout = "layout/main.tpl"
	c.TplName = "user/monitor/delete.tpl"

	monId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	monitor, _ := domain.GetMonitorService().Find(monId)

	c.Data["monitor"] = monitor
}

func (c *DeleteMonitorController) Post() {
	monId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))

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
