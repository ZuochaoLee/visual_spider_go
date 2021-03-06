package controllers

import (
	//"fmt"
	"github.com/astaxie/beego"
	"strconv"
	"strings"
	"visual_spider_go/models"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Conf"] = models.GetConfs()
	c.TplName = "index.tpl"
}

type TestController struct {
	beego.Controller
}

func (c *TestController) Get() {
	url := c.Input().Get("url")
	rule := c.Input().Get("rule")
	fun := c.Input().Get("fun")
	num := c.Input().Get("num")
	c.Data["json"] = map[string]string{"result": models.Test(url, rule, fun, num)}
	c.ServeJSON()
}

type GetRuleController struct {
	beego.Controller
}

func (c *GetRuleController) Get() {
	id := c.Input().Get("id")
	iid, _ := strconv.Atoi(id)
	c.Data["json"] = models.GetRulersById(iid)
	c.ServeJSON()
}

type AddConfController struct {
	beego.Controller
}

func (c *AddConfController) Get() {
	taskname := c.Input().Get("taskname")
	cron := c.Input().Get("cron")
	des := c.Input().Get("des")
	dbtype := c.Input().Get("dbtype")
	dbhost := c.Input().Get("dbhost")
	dbport := c.Input().Get("dbport")
	dbname := c.Input().Get("dbname")
	dbuser := c.Input().Get("dbuser")
	dbpasswd := c.Input().Get("dbpasswd")
	reqtype := c.Input().Get("reqtype")
	rooturl := c.Input().Get("rooturl")
	cookie := c.Input().Get("cookie")
	headerfile := c.Input().Get("headerfile")
	useproxy := c.Input().Get("useproxy")
	texttype := c.Input().Get("texttype")
	postdata := c.Input().Get("postdata")
	pagepre := c.Input().Get("pagepre")
	pagerule := c.Input().Get("pagerule")
	pagefun := c.Input().Get("pagefun")
	pagefour := c.Input().Get("pagefour")
	pagethree := c.Input().Get("pagethree")
	pagetwo := c.Input().Get("pagetwo")
	pageone := c.Input().Get("pageone")
	theardnum := c.Input().Get("theardnum")
	name := c.Input().Get("name")
	rule := c.Input().Get("rule")
	fun := c.Input().Get("fun")
	num := c.Input().Get("num")
	names := strings.Split(name, "|")
	rules := strings.Split(rule, "|")
	funs := strings.Split(fun, "|")
	nums := strings.Split(num, "|")
	if models.AddConf(taskname, cron, des, dbtype, dbhost, dbport, dbname, dbuser, dbpasswd, reqtype, rooturl, cookie, headerfile, useproxy, texttype, postdata, pagepre, pagerule, pagefun, pagefour, pagethree, pagetwo, pageone, theardnum) == 1 {
		id := models.GetIdByName(taskname)
		code := 0
		for i, _ := range names {
			code = models.AddRule(names[i], rules[i], funs[i], nums[i], id)
		}
		c.Data["json"] = map[string]int{"code": code}
	} else {
		c.Data["json"] = map[string]int{"code": 0}
	}
	c.ServeJSON()
}

type UpdateConfController struct {
	beego.Controller
}

func (c *UpdateConfController) Get() {
	id := c.Input().Get("id")
	taskname := c.Input().Get("taskname")
	cron := c.Input().Get("cron")
	des := c.Input().Get("des")
	dbtype := c.Input().Get("dbtype")
	dbhost := c.Input().Get("dbhost")
	dbport := c.Input().Get("dbport")
	dbname := c.Input().Get("dbname")
	dbuser := c.Input().Get("dbuser")
	dbpasswd := c.Input().Get("dbpasswd")
	reqtype := c.Input().Get("reqtype")
	rooturl := c.Input().Get("rooturl")
	cookie := c.Input().Get("cookie")
	headerfile := c.Input().Get("headerfile")
	useproxy := c.Input().Get("useproxy")
	texttype := c.Input().Get("texttype")
	postdata := c.Input().Get("postdata")
	pagepre := c.Input().Get("pagepre")
	pagerule := c.Input().Get("pagerule")
	pagefun := c.Input().Get("pagefun")
	pagefour := c.Input().Get("pagefour")
	pagethree := c.Input().Get("pagethree")
	pagetwo := c.Input().Get("pagetwo")
	pageone := c.Input().Get("pageone")
	theardnum := c.Input().Get("theardnum")
	name := c.Input().Get("name")
	rule := c.Input().Get("rule")
	fun := c.Input().Get("fun")
	ids := c.Input().Get("ids")
	num := c.Input().Get("num")
	names := strings.Split(name, "|")
	rules := strings.Split(rule, "|")
	funs := strings.Split(fun, "|")
	idss := strings.Split(ids, "|")
	nums := strings.Split(num, "|")
	if models.UpdateConf(id, taskname, cron, des, dbtype, dbhost, dbport, dbname, dbuser, dbpasswd, reqtype, rooturl, cookie, headerfile, useproxy, texttype, postdata, pagepre, pagerule, pagefun, pagefour, pagethree, pagetwo, pageone, theardnum) == 1 {
		code := 0
		for i, _ := range idss {
			if idss[i] == "" {
				idi, _ := strconv.Atoi(id)
				code = models.AddRule(names[i], rules[i], funs[i], nums[i], idi)

			} else {
				code = models.UpdateRule(idss[i], names[i], rules[i], funs[i], nums[i])
			}
		}
		c.Data["json"] = map[string]int{"code": code}
	} else {
		c.Data["json"] = map[string]int{"code": 0}
	}
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
	id := c.Input().Get("id")
	name := c.Input().Get("name")
	cron := c.Input().Get("cron")
	c.Data["json"] = map[string]int{"code": models.Start(id, name, cron)}
	c.ServeJSON()
}
