package models

import (
	"fmt"
	"github.com/jakecoffman/cron"
	//"os/exec"
	"strconv"
	"time"
)

var c *cron.Cron

func init() {
	c = cron.New()
}

func InitTask() {
	fmt.Println("start")
	tasks := GetConfsByStatus(1)
	for i := range tasks {
		c.AddFunc(tasks[i].Cron, func() { Run(strconv.Itoa(tasks[i].ID)) }, tasks[i].TaskName)
	}
	c.Start()
	for {
		time.Sleep(100000 * time.Second)
	}
	c.Stop()
}
func AddCron(id, name, cron string) {
	c.AddFunc(cron, func() { Run(id) }, name)
}
func RemoveCron(name string) {
	c.RemoveJob(name)
}
