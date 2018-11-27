package routers

import (
	"lhfproject/beeproject/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
	beego.Router("/testindex", &controllers.TestController{})
	beego.Router("/index1", &controllers.Test2Controller{})
}
