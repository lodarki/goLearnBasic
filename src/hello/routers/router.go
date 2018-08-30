package routers

import (
	"github.com/astaxie/beego"
	"hello/controllers"
)

func init() {
	beego.Include(&controllers.CMSController{})
	beego.Router("/", &controllers.MainController{})
}
