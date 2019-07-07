package User

import "UptimeMonitor/controllers"

type LogoutController struct {
	controllers.BaseController
}

func (c *LogoutController) Post() {
	c.DelSession("userId")
	c.Redirect("/", 302)
}
