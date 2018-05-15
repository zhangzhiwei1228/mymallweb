package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["mymallweb/controllers:AjaxController"] = append(beego.GlobalControllerRouter["mymallweb/controllers:AjaxController"],
		beego.ControllerComments{
			Method: "WelInc",
			Router: `/ajax/WelInc/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
