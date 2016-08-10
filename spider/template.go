package template

import (
	"core/page_processer"
	"core/pipeline"
	"core/spider"
	"fmt"
	"strconv"
	"time"
)

//页面解析模型
func init() {

}
func Run() {
	//配置信息 可以来自数据库
	base := map[string]string{"taskname": "iwjw", "threadnum": "300", "dbtype": "file", "dbhost": "", "dbport": "", "dbdb": "./data11.txt", "dbuser": "", "dbpasswd": ""}
	conf := map[string]string{"rooturl": "http://www.iwjw.com/sale/beijing/p", "texttype": "html", "resqType": "GET", "cookie": "", "headerfile": "", "postdata": "", "proxy": ""}
	page := map[string]string{"pre": "http://www.iwjw.com/", "rule": "#iwjw > div > div.mod-list > div.mod-lists.mb50.clearfix > div.List.mod-border-box.mod-list-shadow > ol > li> div.f-l > h4 > b > a", "fun": "href"}
	rule := map[string]string{"title": "#iwjw > div > div.mod-rent-detail.clearfix.mod-detail > div > div.detail-basic.clearfix > div.detail-infos > div > div.title > h1", "price": "#iwjw > div > div.mod-rent-detail.clearfix.mod-detail > div > div.detail-basic.clearfix > div.detail-infos > div > div.title > div > p > span:nth-child(1) > i"}
	fun := map[string]string{"title": "text", "price": "text"}

	sp := spider.NewSpider(page_processer.NewPageProcesserHtml(conf, page, rule, fun), base["taskname"])
	t1 := time.Now()
	//****自定义的翻页规则开始**************************************************//
	for i := 1; i < 100; i++ {
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
