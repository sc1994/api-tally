package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["api-tally/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["api-tally/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/`,
			AllowHTTPMethods: []string{"post"},
			MethodParams: param.Make(),
			Params: nil})

	beego.GlobalControllerRouter["api-tally/controllers:ScheduleController"] = append(beego.GlobalControllerRouter["api-tally/controllers:ScheduleController"],
		beego.ControllerComments{
			Method: "Get",
			Router: `/`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
