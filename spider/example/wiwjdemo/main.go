//
package main

/*
Packages must be imported:
    "core/common/page"
    "core/spider"
Pckages may be imported:
    "core/pipeline": scawler result persistent;
    "github.com/PuerkitoBio/goquery": html dom parser.
*/
import (
	//"database/sql"
	"fmt"
	// "github.com/hu17889/go_spider/core/common/page"
	// "github.com/hu17889/go_spider/core/pipeline"
	// "github.com/hu17889/go_spider/core/spider"
	"../../core/page_processer"
	"../../core/pipeline"
	"../../core/spider"
	"strconv"
	"strings"
	"time"
)

//页面解析模型

func main() {
	//配置信息 可以来自数据库
	base := map[string]string{"taskname": "iwjw", "threadnum": "30", "dbtype": "file", "dbhost": "", "dbport": "", "dbdb": "./data11.txt", "dbuser": "", "dbpasswd": ""}
	conf := map[string]string{"rooturl": "https://www.iwjw.com/sale/beijing/g1id124|ip5p|/", "texttype": "html", "resqType": "GET", "cookie": "", "headerfile": "", "postdata": "", "proxy": ""}
	page := map[string]string{"pre": "http://www.iwjw.com/", "rule": "#iwjw > div > div.mod-list > div.mod-lists.mb50.clearfix > div.List.mod-border-box.mod-list-shadow > ol > li> div.f-l > h4 > b > a", "fun": "href"}
	rule := map[string]string{"title": "#iwjw > div > div.mod-rent-detail.clearfix.mod-detail > div > div.detail-basic.clearfix > div.detail-infos > div > div.title > h1", "price": "#iwjw > div > div.mod-rent-detail.clearfix.mod-detail > div > div.detail-basic.clearfix > div.detail-infos > div > div.title > div > p > span:nth-child(1) > i"}
	fun := map[string]string{"title": "text", "price": "text"}
	pagerule := map[string]string{"one": "-", "two": "-", "three": "39 41", "four": "1-1-100"}

	pagetype := map[string]string{"one": "list", "two": "list", "three": "list", "four": "num"}
	pageindex := map[string][]string{}

	sp := spider.New(base["taskname"])
	if conf["texttype"] == "html" {
		sp.SetPageProcesser(page_processer.NewPageProcesserHtml(conf, page, rule, fun))
	}
	t1 := time.Now()
	//****自定义的翻页规则开始**************************************************//

	urls := strings.Split(conf["rooturl"], "|")
	l := len(urls)
	var a, b, c int
	for k, v := range pagerule {
		if pagetype[k] == "list" {
			pageindex[k] = strings.Split(v, " ")
		} else {
			pageindex[k] = strings.Split(v, "-")
			t, _ := strconv.Atoi(pageindex[k][0])
			a = t
			tt, _ := strconv.Atoi(pageindex[k][1])
			b = tt
			ttt, _ := strconv.Atoi(pageindex[k][2])
			c = ttt
		}
	}
	fmt.Println(a, b, c, l)
	for i := range pageindex["one"] {
		for ii := range pageindex["two"] {
			for iii := range pageindex["three"] {
				for iiii := a; iiii <= c; iiii = iiii + b {
					url := ""
					fmt.Println(i, ii, iii, iiii)
					if l == 2 {
						url = urls[0] + strconv.Itoa(iiii) + urls[1]
					} else if l == 3 {
						url = urls[0] + pageindex["three"][iii] + urls[1] + strconv.Itoa(iiii) + urls[2]
					} else if l == 4 {
						url = urls[0] + pageindex["two"][ii] + urls[1] + pageindex["three"][iii] + urls[2] + strconv.Itoa(iiii) + urls[3]
					} else if l == 5 {
						url = urls[0] + pageindex["one"][i] + urls[1] + pageindex["two"][ii] + urls[2] + pageindex["three"][iii] + urls[3] + strconv.Itoa(iiii) + urls[4]
					}
					fmt.Println(url)
					sp.AddMyUrl(url, conf["texttype"], "", conf["resqType"], conf["postdata"], conf["proxy"], conf["heardefile"], conf["cookie"])

				}
			}
		}
	}

	//****自定义的翻页规则结束**************************************************//

	if base["dbtype"] == "file" {
		sp.AddPipeline(pipeline.NewPipelineFile(base["dbdb"]))
	} else if base["dbtype"] == "redis" {
		port, _ := strconv.Atoi(base["dbport"])
		db, _ := strconv.Atoi(base["dbdb"])
		sp.AddPipeline(pipeline.NewPipelineRedis(base["dbhost"], port, db, base["dbpasswd"]))
	}
	tn, _ := strconv.Atoi(base["threadnum"])
	sp.
		SetThreadnum(uint(tn)). // Crawl request by three Coroutines
		Run()
	t2 := time.Now()
	fmt.Println(t1, t2)
}
