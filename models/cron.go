package models

import (
	"fmt"
	"github.com/jakecoffman/cron"
	"os/exec"
	"time"
)

var c *cron.Cron

func init() {
	c = cron.New()
}
func InitTask() {
	fmt.Println("start")
	tasks := GetTask()
	for i := 0; i < len(tasks); i++ {
		t := tasks[i].Time
		s := tasks[i].Script
		cm := tasks[i].Command
		n := tasks[i].Name
		c.AddFunc(t, func() { Exec(s, cm) }, n)
	}
	c.Start()
	for {
		time.Sleep(100000 * time.Second)
	}
	c.Stop()
}
func Exec(n, s string) {
	cm := GetConfByName(n)
	cmd := exec.Command(cm, s)
	//buf, err := cmd.Output()
	//fmt.Printf("%s\n%s", buf, err)
}
func AddCron(t, s, cm, n string) {
	c.AddFunc(t, func() { Exec(s, cm) }, n)
}
func RemoveCron(n string) {
	c.RemoveJob(n)
}
