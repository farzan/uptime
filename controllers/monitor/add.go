package User

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/domain"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/validation"
	"strconv"
)

type AddMonitorController struct {
	controllers.BaseController
}

func (c *AddMonitorController) Get() {
	if !domain.GetMonitorService().UserHasReachedMax(c.getUserId()) {
		c.TplName = "user/monitor/no_more_add.tpl"
		return
	}

	c.TplName = "user/monitor/add.tpl"

	monitor := domain.Monitor{
		Thresholds: domain.NewDefaultThreshold(),
	}
	c.Data["monitor"] = monitor
}

func (c *AddMonitorController) Post() {
	c.TplName = "user/monitor/add.tpl"

	ok, _ := c.GetInt("ok", domain.OK)
	warning, _ := c.GetInt("warning", domain.WARNING)
	critical, _ := c.GetInt("critical", domain.CRITICAL)
	interval, _ := c.GetInt("interval", 60)

	monitor := domain.Monitor{
		UserId:   c.getUserId(),
		Title:    c.GetString("title"),
		Method:   domain.GET,
		Url:      c.GetString("url"),
		Interval: interval,
		Thresholds: domain.NewThreshold(
			ok,
			warning,
			critical,
		),
	}

	if !c.isValid(monitor) {
		c.Data["monitor"] = monitor
		return
	}

	domain.GetMonitorService().Add(&monitor)
	domain.GetWorkerService().Start(&monitor)

	c.Redirect("/user", 302)
}

func (c *AddMonitorController) getUserId() int {
	// todo
	return 1
}

func (c *AddMonitorController) isValid(monitor domain.Monitor) bool {
	valid := validation.Validation{}
	isValid := true

	valid.Required(monitor.Title, "title")
	valid.Required(monitor.Method, "method")
	valid.Required(monitor.Interval, "interval")
	valid.Numeric(strconv.Itoa(monitor.Interval), "interval")
	valid.Required(monitor.Url, "url")

	flash := beego.NewFlash()

	if valid.HasErrors() {
		isValid = false
		for _, err := range valid.Errors {
			flash.Error(err.Message)
		}
	}

	if !domain.GetMonitorService().UrlIsUnique(c.getUserId(), monitor.Url) {
		isValid = false
		flash.Error("Url is not unique")
	}

	flash.Store(&c.Controller)

	return isValid
}
