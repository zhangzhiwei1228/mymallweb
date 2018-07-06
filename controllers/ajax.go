package controllers

import (
	"github.com/astaxie/beego"
)

type AjaxController struct {
	BaseController
}

func (c *AjaxController) URLMapping() {
	c.Mapping("WelInc", c.WelInc)
}
// @router /ajax/WelInc/ [get]
func (c *AjaxController) WelInc() {
	beego.Info(string(c.Ctx.Input.RequestBody))
	c.TplName = "ajax/wel_inc.tpl"
}
