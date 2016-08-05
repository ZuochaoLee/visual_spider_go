package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"visual_spider_go/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Lang"] = models.GetConf()
	c.Data["Task"] = models.GetTask()
	c.TplName = "index.tpl"
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	models.InitTask()
	//c.Data["json"] = map[string]int{"code": models.UpdateConf("php", "/lnmp/php/bin/php")}
	c.ServeJSON()
}

type AddTController struct {
	beego.Controller
}

func (c *AddTController) Get() {
	n := c.Input().Get("n")
	t := c.Input().Get("t")
	d := c.Input().Get("d")
	s := c.Input().Get("s")
	o := c.Input().Get("o") //name, script, command, times, des string, status int
	c.Data["json"] = map[string]int{"code": models.AddTask(n, s, o, t, d, 0)}
	c.ServeJSON()
}

type UpdateTController struct {
	beego.Controller
}

func (c *UpdateTController) Get() {
	n := c.Input().Get("n")
	t := c.Input().Get("t")
	d := c.Input().Get("d")
	s := c.Input().Get("s")
	o := c.Input().Get("o") //name, script, command, times, des string, status int
	c.Data["json"] = map[string]int{"code": models.UpdateTask(n, s, o, t, d, 0)}
	c.ServeJSON()
}

type AddCController struct {
	beego.Controller
}

func (c *AddCController) Get() {
	n := c.Input().Get("n")
	o := c.Input().Get("o")
	fmt.Println(n, o)
	c.Data["json"] = map[string]int{"code": models.UpdateConf(n, o)}
	c.ServeJSON()
}

type StopController struct {
	beego.Controller
}

func (c *StopController) Get() {
	name := c.Input().Get("name")
	c.Data["json"] = map[string]int{"code": models.Stop(name)}
	c.ServeJSON()
}

type StartController struct {
	beego.Controller
}

func (c *StartController) Get() {
	name := c.Input().Get("name")
	c.Data["json"] = map[string]int{"code": models.Start(name)}
	c.ServeJSON()
}
