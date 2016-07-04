package routers

import (
	"codes/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/?:page", &controllers.MainController{}, "get:RenderPage")
}
