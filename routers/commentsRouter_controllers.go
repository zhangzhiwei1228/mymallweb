package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["mymallweb/controllers:AccountController"] = append(beego.GlobalControllerRouter["mymallweb/controllers:AccountController"],
		beego.ControllerComments{
			Method: "login",
			Router: `/account/login/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["mymallweb/controllers:AjaxController"] = append(beego.GlobalControllerRouter["mymallweb/controllers:AjaxController"],
		beego.ControllerComments{
			Method: "WelInc",
			Router: `/ajax/WelInc/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
