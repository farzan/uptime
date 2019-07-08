package controllers

type MainController struct {
	BaseController
}

func (c *MainController) Get() {
	if c.GetSession("userId") != nil {
		c.Redirect("/user", 302)
		return
	}

	c.TplName = "index.tpl"
}
