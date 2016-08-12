package models

import (
	"fmt"
	"github.com/jakecoffman/cron"
	//"os/exec"
	"strconv"
	"time"
	"visual_spider_go/spider/template"
)

var c *cron.Cron

func init() {
	c = cron.New()
}

func InitTask() {
	fmt.Println("start")
	tasks := GetConfs()
	for i := range tasks {
		c.AddFunc(tasks[i].Cron, func() { template.Run(strconv.Itoa(tasks[i].ID), db) }, tasks[i].TaskName)
	}
	c.Start()
	for {
		time.Sleep(100000 * time.Second)
	}
	c.Stop()
}
func AddCron(id, name, cron string) {
	println(id, name, cron)
	c.AddFunc(cron, func() { template.Run(id, db) }, name)
}
func RemoveCron(name string) {
	c.RemoveJob(name)
}
