package User

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/domain/entities"
	"UptimeMonitor/domain/services"
	"UptimeMonitor/domain/services/filters"
	"strconv"
)

type ReportController struct {
	controllers.BaseController
}

func(c *ReportController) Get() {
	c.Layout = "layout/main.tpl"
	c.TplName = "user/report.tpl"

	monId, _ := strconv.Atoi(c.Ctx.Input.Param(":id"))
	responses := getResponses(monId)

	c.Data["responses"] = responses
	c.Data["count"] = len(responses)
}

func getResponses(monId int) []entities.Response {
	filter := filters.NewResponseFilter(monId, 100)
	service := services.GetMonitoringResultService()

	return service.FindResponsesByFilter(filter)
}
