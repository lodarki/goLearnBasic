package controllers

import "github.com/astaxie/beego"

type CMSController struct {
	beego.Controller
}

func (c *CMSController) URLMapping() {
	c.Mapping("StaticBlock", c.StaticBlock)
	c.Mapping("AllBlock", c.AllBlock)
}

// @router /staticblock/:key [get]
func (c *CMSController) StaticBlock() {

}

// @router /all/:key [get]
func (c *CMSController) AllBlock() {

}
