package routers

import (
	"github.com/astaxie/beego"
	"socket-client-demo01/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
