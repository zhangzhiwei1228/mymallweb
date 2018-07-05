package controllers

import (
	"github.com/astaxie/beego"
	"log"
)

type AccountController struct {
	beego.Controller
}



//登录页面
func (c *AccountController) LoginTpl() {
	log.Print("loginTpl")
	c.Layout = "inc/login_layout.tpl"
	c.Data["webTitle"] = "zzwrdShop专卖之登录"
	c.TplName = "login/login.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header_login"] = "inc/header_login.tpl"
	c.LayoutSections["footer_login"] = "inc/footer_login.tpl"
}

//登录/注册成功跳转的页面
func (c *AccountController) SuccessTpl() {
	c.Layout = "inc/layout.tpl"
	c.Data["webTitle"] = "zzwrdShop专卖"
	c.TplName = "login/loginSuccess.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header_n"] = "inc/header_n.tpl"
	c.LayoutSections["footer"] = "inc/footer.tpl"
}

//登录提交的页面
func (c *AccountController) DoLogin() {
	jsoninfo := c.GetString("username")
	if jsoninfo == "" {
		c.Ctx.WriteString("username is empty")
		return
	}
}

//注册页面/account/registerTpl
func (c *AccountController) RegisterTpl() {
	log.Print("...........registerTpl.......................")
	c.Layout = "inc/login_layout.tpl"
	c.Data["webTitle"] = "zzwrdShop专卖之注册"
	c.TplName = "login/reg.tpl"
	c.LayoutSections = make(map[string]string)
	c.LayoutSections["header_login"] = "inc/header_login.tpl"
	c.LayoutSections["footer_login"] = "inc/footer_login.tpl"
}

//注册提交的页面
func (c *AccountController) DoRegister() {
	jsoninfo := c.GetString("username")
	if jsoninfo == "" {
		c.Ctx.WriteString("username is empty")
		return
	}
}
