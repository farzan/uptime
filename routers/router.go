package routers

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/controllers/User"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/user", &User.DefaultController{})
    beego.Router("/user/monitor/:id:int/report", &User.ReportController{})
}
