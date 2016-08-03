package routers

import (
	"github.com/astaxie/beego"
	"visual_spider_go/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
