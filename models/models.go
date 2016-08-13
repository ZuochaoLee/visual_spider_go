package models

import (
	"database/sql"
	//"fmt"
	"github.com/astaxie/beego"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

var db *sql.DB

func init() {
	dbs, err := sql.Open("mysql", beego.AppConfig.String("dbuser")+":"+beego.AppConfig.String("dbpassword")+"@tcp("+beego.AppConfig.String("dbhost")+":"+beego.AppConfig.String("dbport")+")/"+beego.AppConfig.String("dbname")+"?charset=utf8")
	db = dbs
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	//defer db.Close()
}
func GetConfs() (conf []Conf) {
	rows, err := db.Query("select * from conf ")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	ct := Conf{}
	for rows.Next() {
		err := rows.Scan(&ct.ID, &ct.TaskName, &ct.Theardnum, &ct.Cron, &ct.Des, &ct.Dbtyble, &ct.Dbhost, &ct.Dbport, &ct.Dbname, &ct.Dbuser, &ct.Dbpasswd, &ct.ReqType, &ct.RootUrl, &ct.Cookie, &ct.HeaderFile, &ct.UseProxy, &ct.TextType, &ct.PostData, &ct.Status, &ct.PagePre, &ct.PageRule, &ct.PageFun, &ct.PageFour, &ct.PageThree, &ct.PageTwo, &ct.PageOne)
		if err != nil {
			log.Fatal(err)
		}
		conf = append(conf, ct)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}
func GetConfsByStatus(s int) (conf []Conf) {
	rows, err := db.Query("select * from conf where status=" + strconv.Itoa(s) + ";")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	ct := Conf{}
	for rows.Next() {
		err := rows.Scan(&ct.ID, &ct.TaskName, &ct.Theardnum, &ct.Cron, &ct.Des, &ct.Dbtyble, &ct.Dbhost, &ct.Dbport, &ct.Dbname, &ct.Dbuser, &ct.Dbpasswd, &ct.ReqType, &ct.RootUrl, &ct.Cookie, &ct.HeaderFile, &ct.UseProxy, &ct.TextType, &ct.PostData, &ct.Status, &ct.PagePre, &ct.PageRule, &ct.PageFun, &ct.PageFour, &ct.PageThree, &ct.PageTwo, &ct.PageOne)
		if err != nil {
			log.Fatal(err)
		}
		conf = append(conf, ct)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}
func GetConfById(id int) (conf Conf) {
	sid := strconv.Itoa(id)
	rows, err := db.Query("select * from conf where id=" + sid)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&conf.ID, &conf.TaskName, &conf.Theardnum, &conf.Cron, &conf.Des, &conf.Dbtyble, &conf.Dbhost, &conf.Dbport, &conf.Dbname, &conf.Dbuser, &conf.Dbpasswd, &conf.ReqType, &conf.RootUrl, &conf.Cookie, &conf.HeaderFile, &conf.UseProxy, &conf.TextType, &conf.PostData, &conf.Status, &conf.PagePre, &conf.PageRule, &conf.PageFun, &conf.PageFour, &conf.PageThree, &conf.PageTwo, &conf.PageOne)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}
func GetRulersById(id int) (ruler []Ruler) {
	sid := strconv.Itoa(id)
	rows, err := db.Query("select * from rule where taskid=" + sid)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	ct := Ruler{}
	for rows.Next() {
		err := rows.Scan(&ct.ID, &ct.TaskId, &ct.Name, &ct.Rule, &ct.Fun)
		if err != nil {
			log.Fatal(err)
		}
		ruler = append(ruler, ct)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}
func AddConf(taskname, cron, des, dbtype, dbhost, dbport, dbname, dbuser, dbpasswd, reqtype, rooturl, cookie, headerfile, useproxy, texttype, postdata, pagepre, pagerule, pagefun, pagefour, pagethree, pagetwo, pageone, threadnum string) (code int) {
	var sql = "INSERT INTO conf(taskname, cron, des, dbtype, dbhost, dbport, dbname, dbuser, dbpasswd, reqtype, rooturl, cookie, headerfile, useproxy, texttype, postdata, pagepre, pagerule, pagefun,pagefour,pagethree,pagetwo,pageone,theardnum,status) VALUES('" + taskname + "', '" + cron + "', '" + des + "', '" + dbtype + "','" + dbhost + "','" + dbport + "','" + dbname + "','" + dbuser + "','" + dbpasswd + "','" + reqtype + "','" + rooturl + "','" + cookie + "','" + headerfile + "','" + useproxy + "','" + texttype + "','" + postdata + "','" + pagepre + "','" + pagerule + "','" + pagepre + "','" + pagefour + "','" + pagethree + "','" + pagetwo + "','" + pageone + "', " + threadnum + ", 0);"
	//println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		code = 0
	} else {
		code = 1
	}
	return
}

func GetIdByName(taskname string) (id int) {
	rows, err := db.Query("select id from conf where taskname='" + taskname + "';")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			log.Fatal(err)
		}
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}
func AddRules(name, rule, fun string, taskid int) (code int) {

	var sql = "INSERT INTO rule(taskid,name, rule, fun) VALUES(" + strconv.Itoa(taskid) + ",'" + name + "', '" + rule + "', '" + fun + ");"
	//println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		code = 0
	} else {
		code = 1
	}

	return
}

// func UpdateTask(name, script, command, times, des string, status int) (code int) {
// 	var sql = "update task set name='" + name + "', script='" + script + "',command='" + command + "',times='" + times + "',status=" + strconv.Itoa(status) + ",des='" + des + "' where name='" + name + "';"
// 	_, err := db.Exec(sql)
// 	if err != nil {
// 		code = 0
// 	} else {
// 		code = 1
// 	}
// 	return
// }
// func UpdateConf(name, script string) (code int) {
// 	var sql = "select count(*) from conf where name='" + name + "';" //'" + name + "';"
// 	rows, _ := db.Query(sql)
// 	defer rows.Close()
// 	var num int
// 	for rows.Next() {
// 		err := rows.Scan(&num)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		if num == 0 {
// 			sql = "INSERT INTO conf (name, script) VALUES ('" + name + "','" + script + "');"
// 			_, err := db.Exec(sql)
// 			if err != nil {
// 				code = 0
// 			} else {
// 				code = 1
// 			}
// 		} else {
// 			sql = "update conf set script='" + script + "' where name='" + name + "';"
// 			_, err := db.Exec(sql)
// 			if err != nil {
// 				code = 0
// 			} else {
// 				code = 1
// 			}
// 		}
// 	}
// 	return
// }
func Stop(name string) (code int) {
	var sql = "update conf set status=0 where taskname='" + name + "';"
	_, err := db.Exec(sql)
	if err != nil {
		code = 0
	} else {
		code = 1
		RemoveCron(name)
	}
	return
}
func Start(id, name, cron string) (code int) {
	var sql = "update conf set status=1 where id=" + id + ";"
	_, err := db.Exec(sql)
	if err != nil {
		code = 0
	} else {
		code = 1
		AddCron(id, name, cron)
	}
	return
}
