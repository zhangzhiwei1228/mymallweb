package controllers

import (
	"github.com/astaxie/beego"
)

type AjaxController struct {
	beego.Controller
}

func (c *AjaxController) URLMapping() {
	c.Mapping("WelInc", c.WelInc)
}
// @router /ajax/WelInc/ [get]
func (c *AjaxController) WelInc() {
	c.TplName = "ajax/wel_inc.tpl"
}
