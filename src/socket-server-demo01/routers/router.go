package routers

import (
	"github.com/astaxie/beego"
	"socket-server-demo01/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
