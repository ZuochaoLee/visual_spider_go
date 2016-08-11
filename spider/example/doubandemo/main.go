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
	"fmt"

	// "github.com/hu17889/go_spider/core/common/page"
	// "github.com/hu17889/go_spider/core/pipeline"
	// "github.com/hu17889/go_spider/core/spider"
	"../../core/page_processer"
	"../../core/pipeline"
	"../../core/spider"
	"strconv"
	"time"
)

//页面解析模型

func main() {
	//配置信息 可以来自数据库
	base := map[string]string{"taskname": "shuimu", "threadnum": "3", "dbtype": "file", "dbhost": "", "dbport": "", "dbdb": "./datashuimu.txt", "dbuser": "", "dbpasswd": ""}
	conf := map[string]string{"rooturl": "http://www.newsmth.net/nForum/#!board/HouseRent?p=", "texttype": "html", "resqType": "GET", "cookie": "", "headerfile": "", "postdata": "", "proxy": ""}
	page := map[string]string{"pre": "http://www.newsmth.net/", "rule": "#body > div.b-content > table > tbody > tr > td.title_9.bg-odd > a", "fun": "href"}
	rule := map[string]string{"body": "#body > div.b-content.corner > div:nth-child(2) > table > tbody > tr.a-body > td.a-content > p", "pic": ".topic-figure.cc > img"}
	fun := map[string]string{"body": "text", "pic": "src"}

	sp := spider.NewSpider(page_processer.NewPageProcesserHtml(conf, page, rule, fun), base["taskname"])
	t1 := time.Now()
	//****自定义的翻页规则开始**************************************************//
	for i := 1; i < 2; i++ {
		sp.AddMyUrl(conf["rooturl"]+strconv.Itoa(i), conf["texttype"], "", conf["resqType"], conf["postdata"], conf["proxy"], conf["heardefile"], conf["cookie"])
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
