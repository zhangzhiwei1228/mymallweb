package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	BaseController
}

func (c *DefaultController) Get() {
	beego.Info(string(c.Ctx.Input.RequestBody))
	c.Layout = "inc/layout.tpl"
	c.Data["webTitle"] = "zzwrdShop专卖"
	c.TplName = "index.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header_n"] = "inc/header_n.tpl"
	c.LayoutSections["footer"] = "inc/footer.tpl"
}
