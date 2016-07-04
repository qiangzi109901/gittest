package routers

import (
	"github.com/astaxie/beego"
	"testhttp/controllers"
)

func init() {
	beego.Router("/?:page", &controllers.MainController{}, "get:RenderPage")
}
