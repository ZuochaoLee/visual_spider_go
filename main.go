package main

import (
	"github.com/astaxie/beego"
	"visual_spider_go/models"
	_ "visual_spider_go/routers"
)

func main() {
	go models.InitTask()
	beego.Run()
}
