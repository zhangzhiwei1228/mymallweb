package controllers

import (
	"github.com/astaxie/beego"
	"log"
	"encoding/json"
	_ "time"

	"github.com/astaxie/beego/orm"
	"mymallweb/models"
	"mymallweb/library/helper"
	"encoding/gob"
	"strings"
	"time"
	"strconv"
)

type AccountController struct {
	BaseController
}
func init()  {
	var user models.User
	gob.Register(user)
}

//登录页面
func (c *AccountController) LoginTpl() {
	SessonInfo := c.GetSession("USER_LOGIN_VALUE")
	if SessonInfo != nil {
		c.Redirect("/",302 )
	}
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
	c.LayoutSections["header_n"] = "inc/header_success.tpl"
	c.LayoutSections["footer"] = "inc/footer.tpl"
}

//登录提交的页面
func (c *AccountController) DoLogin() {
	SessonInfo := c.GetSession("USER_LOGIN_VALUE")
	if SessonInfo != nil {
		c.Redirect("/",302 )
	}
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
	o := orm.NewOrm()
	beego.Info("开始查询")
	err = o.Raw("SELECT id,nickname,password,salt,role,username FROM `scn_user` WHERE username = ? OR email = ? OR mobile = ? and role='member' limit 1", account_name).QueryRow(&user)
	beego.Info(err)
	if err != nil {
		c.SetJson(1, nil, "此用户不存在，请确认后重试")
		return
	}
	pwdMd5 := helper.PwdEncrypt(password,user.Salt)
	if pwdMd5 != user.Password {
		c.SetJson(1, nil, "用户密码不正确")
		return
	}
	sessionData := map[string]interface{}{"username": user.Username,"nickname":user.Nickname,"role":user.Role}
	jsonSessionData, err := json.Marshal(sessionData)
	if err != nil {
		beego.Error("json.Marshal failed:", err)
		return
	}
	beego.Info(string(jsonSessionData))
	c.SetSession("USER_LOGIN_VALUE",string(jsonSessionData))
	c.SetSession("USER_KEY_USERID",user.Id)
	res := map[string]string{} //声明空数组
	//os.Exit(1) //调试用
	res["nickname"] = user.Nickname
	res["username"] = user.Username
	c.SetJson(0, res, "登录成功")
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
//检查是否登录
func (c *AccountController) CheckInfo()  {
	var code int
	SessonInfo := c.GetSession("USER_LOGIN_VALUE")
	var SessionInfoStr string
	beego.Info(SessonInfo)
	var accountInfo map[string]interface{}
	if SessonInfo != nil {
		switch v := SessonInfo.(type) {
		case string:
			SessionInfoStr = string(v)
		}
		json.Unmarshal([]byte(SessionInfoStr), &accountInfo)
		beego.Info(accountInfo)
		code = 0
	} else {
		code = 1
	}
	c.SetJson(code,accountInfo,"")
	return
}
func (c *AccountController)  LogOut(){
	SessonInfo := c.GetSession("USER_LOGIN_VALUE")
	if SessonInfo != nil {
		c.DelSession("USER_LOGIN_VALUE")
	}
	c.SetJson(0,"","")
	return
}

func (c *AccountController) GetVerfiCode() {
	c.SetJson(0, 1, "获取手机验证码成功")
	return
	beego.Info(string(c.Ctx.Input.RequestBody))
	postData := map[string]string{"mobile": "", "captchaVal": "","captchaKey":"","is_captcha":""}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &postData)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	captcha := postData["is_captcha"]
	beego.Info(captcha)
	if captcha != "true" {
		c.SetJson(1, nil, "图形验证码输入错误")
		return
	}
	captchaVal := postData["captchaVal"]
	captchaKey := postData["captchaKey"]
	beego.Info(captchaKey)
	beego.Info(captchaVal)
	if !helper.ValidateCaptcha(captchaKey,captchaVal,false) {
		c.SetJson(1, nil, "输入的图形验证码错误")
		return
	}
	mobile := postData["mobile"]
	if !helper.ValidateMobile(mobile) {
		c.SetJson(1, nil, "手机号格式错误")
		return
	}

	//获取真实ip
	ip := c.Ctx.Request.Header.Get("X-Forwarded-For")
	inputIp := c.Ctx.Input.IP()
	if ip != inputIp {
		c.SetJson(1, nil, "用户ip存在异常")
		return
	}
	beego.Info("通过X-Forwarded-For获取的ip地址 ：" +ip)
	beego.Info("通过c.Ctx.Input.IP获取的ip地址 ：" +inputIp)
	//查询数据库 ip or mobile having count < 4 最新时间小于当前时间
	var ipInt = helper.IpToInt(ip)
	timeUnix:=time.Now().Unix()
	SelectData := []string{mobile, strconv.Itoa(ipInt), strconv.FormatInt(timeUnix,10)}
	beego.Info(SelectData)
	var captchaModel models.Captcha
	o := orm.NewOrm()
	beego.Info("开始查询")
	err = o.Raw("SELECT id FROM `scn_captcha` WHERE (mobile = ? or ip = ?) and expiration_date >= ? limit 1", SelectData).QueryRow(&captchaModel)
	beego.Info(err)
	if err == nil {
		c.SetJson(1, nil, "不要重复获取")
		return
	}

	//随机生成6位数字
	VerfiCode := helper.RandCaptchaCode()
	beego.Info("手机号为："+mobile+" ，ip地址为："+inputIp+" ，生成的验证码为："+VerfiCode)

	//发送短信
	smsMsg := "您的注册验证码为："+VerfiCode+"，祝您购物愉快。"
	smsUrl :=smsMsg;
	res := helper.HttpGet(smsUrl) //20180810142716,0 8900810142716072000
	beego.Info(res)
	//var SmsInfo map[string]string{}
	SmsInfo := strings.Split(res, ",")
	beego.Info("发送的手机验证码返回的信息：" + res)
	var smsStatus int
	smsStatus, _ =strconv.Atoi(SmsInfo[1])
	if smsStatus != 0 {
		c.SetJson(1, nil, "获取手机验证码失败，请重新获取")
		return
	}
	var captchaInsterData  models.Captcha

	var CaptchaModelValue int
	var VerfiCodeInt int
	var create_date int64
	var result int64
	var ExpirationDate int64

	create_date = time.Now().Unix()
	mm, _ := time.ParseDuration("15m") //过期时间大于十五分钟
	ExpirationDate = time.Now().Add(mm).Unix()
	CaptchaModelValue, _ = strconv.Atoi(captchaVal) //输入的图形验证码值
	VerfiCodeInt, _ = strconv.Atoi(VerfiCode) //发送的手机验证码
	captchaInsterData.Mobile = mobile
	captchaInsterData.Captcha = CaptchaModelValue
	captchaInsterData.CaptchaKey = captchaKey
	captchaInsterData.VerfiCode = VerfiCodeInt
	captchaInsterData.CreateDate = create_date
	captchaInsterData.ExpirationDate = ExpirationDate
	captchaInsterData.Ip = ipInt
	models.AddCaptcha(&captchaInsterData)
	c.SetSession("USER_MOBILE_VERFICODE",VerfiCodeInt)
	c.SetJson(0, result, "获取手机验证码成功")
	return

}

func (c *AccountController) TestFun()  {
	ip := c.Ctx.Request.Header.Get("HTTP_X_REAL_IP")
	beego.Info(ip)
	beego.Info(c.Ctx.Input.IP())
	return
}
