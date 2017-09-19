package routers

import (
	"MaintainServer/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/login", &controllers.MainController{},"*:Login")
    beego.Router("/logout", &controllers.MainController{},"get:Logout")
    beego.Router("/add", &controllers.MainController{},"*:Add")
    beego.Router("/edit/:id:int", &controllers.MainController{},"*:Edit")
    beego.Router("/del/:id:int", &controllers.MainController{},"*:Del")
    beego.Router("/", &controllers.MainController{},"get:List")
    beego.Router("/api", &controllers.ApiController{},"get:Api")
}
