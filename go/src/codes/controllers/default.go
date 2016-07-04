package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

func (c *MainController) RenderPage() {

	c.Layout = "layout/home.html"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["Header"] = "control/home/head.html"
	c.LayoutSections["Sidebar"] = "control/home/sidebar.html"
	c.LayoutSections["Footer"] = "control/home/footer.html"

	page := c.Ctx.Input.Param(":page")
	if page == "" {
		page = "index"
	}
	c.TplName = "screen/home/" + page + ".html"

	c.Data["LayoutClass"] = "teal"
}
