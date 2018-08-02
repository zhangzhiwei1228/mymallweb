package controllers

import (
	"github.com/astaxie/beego"
)

type DefaultController struct {
	BaseController
}

//var user interface{}
func (c *DefaultController) Get() {
	c.Layout = "inc/layout.tpl"
	c.Data["webTitle"] = "zzwrdShop专卖"
	c.TplName = "index.tpl"
	sessionInfo := map[string]interface{}{"username":"","nickname":"","role":""}
	var Info interface{}
	//var ob models.Object
	Info = c.GetSession("USER_LOGIN_VALUE")
	//err := json.Unmarshal([]byte(Info), &sessionInfo)
	beego.Info(Info)
	beego.Info(sessionInfo)
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header_n"] = "inc/header_n.tpl"
	c.LayoutSections["footer"] = "inc/footer.tpl"
}
