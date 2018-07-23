package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"encoding/json"
	_ "time"

	"github.com/astaxie/beego/orm"
	"mymallweb/models"
	"mymallweb/library/helper"
)

type AccountController struct {
	BaseController
}


//登录页面
func (c *AccountController) LoginTpl() {
	c.Layout = "inc/login_layout.tpl"
	//c.Data["xsrf_token"] = c.XSRFToken()
	c.Data["webTitle"] = "登录"
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
	//哈希校验成功后 更新 auth_key
	beego.Info(string(c.Ctx.Input.RequestBody))
	postData := map[string]string{"username": "", "password": ""}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &postData)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	userName := postData["username"]
	password := postData["password"]
	if userName == "" || password == "" {
		c.SetJson(1, nil, "用户名或密码不能为空")
		return
	}
	account_name := []string{userName, userName, userName}
	var user models.User
	orm.Debug = true
	o := orm.NewOrm()
	beego.Info("开始查询")
	err = o.Raw("SELECT nickname,password,salt,role FROM `scn_user` WHERE username = ? OR email = ? OR mobile = ? and role='member' limit 1", account_name).QueryRow(&user)
	if err != nil {
		c.SetJson(1, nil, "此用户不存在，请确认后重试")
		return
	}
	pwdMd5 := helper.PwdEncrypt(password,user.Salt)
	if pwdMd5 != user.Password {
		c.SetJson(1, nil, "用户密码不正确")
		return
	}
	data := []string{user.Nickname}
	c.SetJson(0, data, "登录成功")
	return
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
	beego.Info(string(c.Ctx.Input.RequestBody))
	jsoninfo := c.GetString("username")
	if jsoninfo == "" {
		c.Ctx.WriteString("username is empty")
		return
	}
}
