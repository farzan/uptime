package routers

import (
	"UptimeMonitor/controllers"
	"UptimeMonitor/controllers/User"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    beego.Router("/user/register", &User.RegisterController{})
    beego.Router("/user/login", &User.LoginController{})
    beego.Router("/user/logout", &User.LogoutController{})

    beego.Router("/user", &monitor.DefaultController{})
    beego.Router("/user/monitor/:id:int/report", &monitor.ReportController{})
    beego.Router("/user/monitor/add", &monitor.AddMonitorController{})
    beego.Router("/user/monitor/:id:int/delete", &monitor.DeleteMonitorController{})
}
