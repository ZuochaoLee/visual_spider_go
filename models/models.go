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
	dbs, err := sql.Open("mysql", beego.AppConfig.String("dbuser")+":"+beego.AppConfig.String("dbpassword")+"@tcp("+beego.AppConfig.String("dbhost")+":"+beego.AppConfig.String("dbpost")+")/"+beego.AppConfig.String("dbname")+"?charset=utf8")
	db = dbs
	if err != nil {
		log.Fatalf("Open database error: %s\n", err)
	}
	//defer db.Close()
}
func GetConf() (conf []Lang) {
	rows, err := db.Query("select name, script from conf ")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	ct := Lang{}
	conf = []Lang{}
	for rows.Next() {
		err := rows.Scan(&ct.Name, &ct.Script)
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
func GetTask() (task []Task) {
	rows, err := db.Query("select name, script,command,times,status,des from task ")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	tt := Task{}
	task = []Task{}
	for rows.Next() {
		err := rows.Scan(&tt.Name, &tt.Script, &tt.Command, &tt.Time, &tt.Status, &tt.Des)
		if err != nil {
			log.Fatal(err)
		}
		task = append(task, tt)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return
}
func GetTaskByName(name string) (task Task) {
	rows, err := db.Query("select name, script,command,times,status,des from task where name='" + name + "';")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	task = Task{}
	for rows.Next() {
		err := rows.Scan(&task.Name, &task.Script, &task.Command, &task.Time, &task.Status, &task.Des)
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
func GetConfByName(name string) (command string) {
	rows, err := db.Query("select script from conf where name='" + name + "';")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&command)
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
func AddTask(name, script, command, times, des string, status int) (code int) {
	var sql = "INSERT INTO task(name, script,command,times,status,des) VALUES('" + name + "', '" + script + "', '" + command + "', '" + times + "', " + strconv.Itoa(status) + ", '" + des + "');"
	println(sql)
	_, err := db.Exec(sql)
	if err != nil {
		code = 0
	} else {
		code = 1
	}
	return
}
func UpdateTask(name, script, command, times, des string, status int) (code int) {
	var sql = "update task set name='" + name + "', script='" + script + "',command='" + command + "',times='" + times + "',status=" + strconv.Itoa(status) + ",des='" + des + "' where name='" + name + "';"
	_, err := db.Exec(sql)
	if err != nil {
		code = 0
	} else {
		code = 1
	}
	return
}
func UpdateConf(name, script string) (code int) {
	var sql = "select count(*) from conf where name='" + name + "';" //'" + name + "';"
	rows, _ := db.Query(sql)
	defer rows.Close()
	var num int
	for rows.Next() {
		err := rows.Scan(&num)
		if err != nil {
			log.Fatal(err)
		}
		if num == 0 {
			sql = "INSERT INTO conf (name, script) VALUES ('" + name + "','" + script + "');"
			_, err := db.Exec(sql)
			if err != nil {
				code = 0
			} else {
				code = 1
			}
		} else {
			sql = "update conf set script='" + script + "' where name='" + name + "';"
			_, err := db.Exec(sql)
			if err != nil {
				code = 0
			} else {
				code = 1
			}
		}
	}
	return
}
func Stop(name string) (code int) {
	var sql = "update task set status=0 where name='" + name + "';"
	_, err := db.Exec(sql)
	if err != nil {
		code = 0
		RemoveCron(name)
	} else {
		code = 1
	}
	return
}
func Start(name string) (code int) {
	var sql = "update task set status=1 where name='" + name + "';"
	_, err := db.Exec(sql)
	if err != nil {
		code = 0
		task := GetTaskByName(name)
		AddCron(task.Time, task.Script, task.Command, task.Name)
	} else {
		code = 1
	}
	return
}
