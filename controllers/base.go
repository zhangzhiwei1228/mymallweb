package controllers

import (
	_"runtime"
	"github.com/astaxie/beego"
	_"github.com/astaxie/beego/orm"
	_ "github.com/astaxie/beego/session/redis"

)
//基类
type BaseController struct {
	beego.Controller
}
func init()  {

}
func (this *BaseController) SetJson(code int, data interface{}, Msg string) {
	if code == 0 && Msg == "" {
		Msg = "sucess"
	}
	this.Data["json"] = map[string]interface{}{"code": code, "msg": Msg, "data": data}
	this.ServeJSON()
}
