package controllers

import (
	"github.com/astaxie/beego"
)

type AccountController struct {
	beego.Controller
}



//登录页面
func (c *AccountController) loginTpl() {
	c.Layout = "inc/login_layout.tpl"
	c.Data["webTitle"] = "zzwrdShop专卖--登录"
	c.TplName = "login/login.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header_login"] = "inc/header_login.tpl"
	c.LayoutSections["footer_login"] = "inc/footer_login.tpl"
}

//登录成功跳转的页面
func (c *AccountController) loginAfter() {
	c.Layout = "inc/login_layout.tpl"
	c.Data["webTitle"] = "zzwrdShop专卖--登录"
	c.TplName = "login/login.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header_login"] = "inc/header_login.tpl"
	c.LayoutSections["footer_login"] = "inc/footer_login.tpl"
}

//登录提交的页面
func (c *AccountController) doLogin() {
	c.TplName = "ajax/wel_inc.tpl"
}


//注册页面
func (c *AccountController) registerTpl() {
	c.Layout = "inc/login_layout.tpl"
	c.Data["webTitle"] = "zzwrdShop专卖--登录"
	c.TplName = "login/login.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header_login"] = "inc/header_login.tpl"
	c.LayoutSections["footer_login"] = "inc/footer_login.tpl"
}

//注册提交的页面
func (c *AccountController) doRegister() {

}
