package routers

import (
	"github.com/astaxie/beego"
	"visual_spider_go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/getrule", &controllers.GetRuleController{})
	beego.Router("/addconf", &controllers.AddConfController{})
	beego.Router("/test", &controllers.TestController{})
	beego.Router("/updateconf", &controllers.UpdateConfController{})
	beego.Router("/stop", &controllers.StopController{})
	beego.Router("/start", &controllers.StartController{})
}
