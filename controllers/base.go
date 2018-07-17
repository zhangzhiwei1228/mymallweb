package controllers

import (
	_"runtime"
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"

	"encoding/json"
)
//基类
type BaseController struct {
	beego.Controller
}

func Init(c *BaseController)  {

}
func (c *BaseController) CheckXSRFCookie() bool {
	postData := map[string]string{ "_xsrf": ""}
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &postData)
	xsrfToken := c.XSRFToken()
	if  err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return false
	}
	if  postData["_xsrf"] == "" {
		c.SetJson(1, nil, "数据缺少")
		return false
	}
	if xsrfToken != postData["_xsrf"] {
		c.SetJson(1, nil, "xsrf有错")
		return false
	}
	return true
}

func (this *BaseController) SetJson(code int, data interface{}, Msg string) {
	if code == 0 && Msg == "" {
		Msg = "sucess"
	}
	this.Data["json"] = map[string]interface{}{"code": code, "msg": Msg, "data": data}
	this.ServeJSON()
}
