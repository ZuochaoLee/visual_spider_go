package controllers

import (
	"fmt"
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

type GetRuleController struct {
	beego.Controller
}

func (c *GetRuleController) Get() {
	id := c.Input().Get("id")
	iid, _ := strconv.Atoi(id)
	c.Data["json"] = models.GetRulersById(iid)
	c.ServeJSON()
}

// type TestController struct {
// 	beego.Controller
// }

// func (c *TestController) Get() {
// 	template.Run()
// 	//c.Data["json"] = map[string]int{"code": models.UpdateConf("php", "/lnmp/php/bin/php")}
// 	c.ServeJSON()
// }

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
	names := strings.Split(name, "|")
	rules := strings.Split(rule, "|")
	funs := strings.Split(fun, "|")
	fmt.Println(names, rules, funs)
	if models.AddConf(taskname, cron, des, dbtype, dbhost, dbport, dbname, dbuser, dbpasswd, reqtype, rooturl, cookie, headerfile, useproxy, texttype, postdata, pagepre, pagerule, pagefun, pagefour, pagethree, pagetwo, pageone, theardnum) == 1 {
		id := models.GetIdByName(taskname)
		fmt.Println(id)
		code := 0
		for i, _ := range names {
			code = models.AddRule(names[i], rules[i], funs[i], id)
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
	names := strings.Split(name, "|")
	rules := strings.Split(rule, "|")
	funs := strings.Split(fun, "|")
	idss := strings.Split(ids, "|")
	fmt.Println(rules, idss)
	if models.UpdateConf(id, taskname, cron, des, dbtype, dbhost, dbport, dbname, dbuser, dbpasswd, reqtype, rooturl, cookie, headerfile, useproxy, texttype, postdata, pagepre, pagerule, pagefun, pagefour, pagethree, pagetwo, pageone, theardnum) == 1 {
		fmt.Println("123")
		code := 0
		for i, _ := range idss {
			fmt.Println(idss[i], name[i])
			if idss[i] == "" {
				idi, _ := strconv.Atoi(id)
				code = models.AddRule(names[i], rules[i], funs[i], idi)

			} else {
				code = models.UpdateRule(idss[i], names[i], rules[i], funs[i])
			}
		}
		c.Data["json"] = map[string]int{"code": code}
	} else {
		c.Data["json"] = map[string]int{"code": 0}
	}
	c.ServeJSON()
}

// type UpdateTController struct {
// 	beego.Controller
// }

// func (c *UpdateTController) Get() {
// 	n := c.Input().Get("n")
// 	t := c.Input().Get("t")
// 	d := c.Input().Get("d")
// 	s := c.Input().Get("s")
// 	o := c.Input().Get("o") //name, script, command, times, des string, status int
// 	c.Data["json"] = map[string]int{"code": models.UpdateTask(n, s, o, t, d, 0)}
// 	c.ServeJSON()
// }

// type AddCController struct {
// 	beego.Controller
// }

// func (c *AddCController) Get() {
// 	n := c.Input().Get("n")
// 	o := c.Input().Get("o")
// 	fmt.Println(n, o)
// 	c.Data["json"] = map[string]int{"code": models.UpdateConf(n, o)}
// 	c.ServeJSON()
// }

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
