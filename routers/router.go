package routers

import (
	"mymallweb/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})

    //自动匹配
	//beego.AutoRouter(&controllers.AjaxController{})

	//注解路由
	beego.Include(&controllers.AjaxController{})
	//beego.Include(&controllers.AccountController{})
	//beego.Router("/account/loginTpl", &controllers.AccountController{}, "get:loginTpl")
}
