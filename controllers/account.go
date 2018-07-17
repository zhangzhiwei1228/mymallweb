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
	password := postData["username"]
	userName := postData["password"]
	if userName == "" || password == "" {
		c.SetJson(1, nil, "用户名或密码不存在")
		return
	}
	var user models.User
	orm.Debug = true
	o := orm.NewOrm()
	err = o.Raw("SELECT * FROM `act_user` WHERE username = ? OR email = ? OR mobile = ?", userName).QueryRow(&user)
	beego.Info(err)
	beego.Info(user)
	c.SetJson(0, user, "")
	//salt := "123456"
	pwdMd5 := helper.Util{}
	beego.Info(pwdMd5)
	//err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	/*if err != nil {
		c.SetJson(1, nil, "用户名或密码错误")
		return
	} else {
		if user.AuthKey == "" {
			userAuth := common.Md5String(user.Username + common.GetString(time.Now().Unix()))
			user.AuthKey = userAuth
			models.UpdateUserById(&user)
		}
		user.PasswordHash = ""
		c.SetJson(0, user, "")
		return
	}*/
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
