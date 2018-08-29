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
	//这里指定了 index.tpl，如果用户不设置该参数，那么默认会去到模板目录的 Controller/<方法名>.tpl 查找，
	// 例如上面的方法会去 maincontroller/get.tpl (文件、文件夹必须小写)。

	//用户设置了模板之后系统会自动的调用 Render 函数（这个函数是在 beego.Controller 中实现的），所以无需用户自己来调用渲染。
	//当然也可以不使用模版，直接用 this.Ctx.WriteString 输出字符串
	c.TplName = "index.tpl"
}
