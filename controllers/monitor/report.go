package monitor

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/domain"
	"strconv"
)

type ReportController struct {
	controllers.BaseController
}

func(c *ReportController) Get() {
	c.TplName = "user/monitor/report.tpl"

	monId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	monitor, _ := domain.GetMonitorService().Find(monId)

	if !c.IsAuthorized(monitor.UserId) {
		return
	}

	responses := getResponses(monId)

	c.Data["responses"] = responses
	c.Data["count"] = len(responses)
}

func getResponses(monId int) []domain.Response {
	filter := domain.NewResponseFilter(monId, 100)
	service := domain.GetMonitoringResultService()

	return service.FindResponsesByFilter(filter)
}
