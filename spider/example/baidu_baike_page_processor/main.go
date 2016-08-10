//
package main

import (
	"fmt"
	"github.com/hu17889/go_spider/core/common/page"
	"github.com/hu17889/go_spider/core/common/request"
	"github.com/hu17889/go_spider/core/spider"
	"strings"
)

type MyPageProcesser struct {
}

func NewMyPageProcesser() *MyPageProcesser {
	return &MyPageProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package goquery (http://godoc.org/github.com/PuerkitoBio/goquery) is used to parse html.
func (this *MyPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	rule := map[string]string{"title": ".house-tit", "price": ".font-price"}
	fun := map[string]string{"title": "text", "price": "text"}
	result := map[string]string{"title": "", "price": ""}
	query := p.GetHtmlParser()
	for k, v := range rule {
		if fun[k] == "text" {
			result[k] = query.Find(v).Text()
		} else {
			result[k], _ = query.Find(v).Attr(fun[k])
		}
		result[k] = strings.Trim(result[k], " \t\n")
		// if result[k] == "" {
		// 	p.SetSkip(true)
		// }
		// the entity we want to save by Pipeline
		p.AddField(k, result[k])
	}

	// name := query.Find(".lemmaTitleH1").Text()
	// name = strings.Trim(name, " \t\n")
	// p.AddField("name", name)
	// summary := query.Find(".card-summary-content .para").Text()
	// summary = strings.Trim(summary, " \t\n")

	// // the entity we want to save by Pipeline

	// p.AddField("summary", summary)
}

func (this *MyPageProcesser) Finish() {
	fmt.Printf("TODO:before end spider \r\n")
}

func main() {
	// spider input:
	//  PageProcesser ;
	//  task name used in Pipeline for record;
	sp := spider.NewSpider(NewMyPageProcesser(), "TaskName")
	// GetWithParams Params:
	//  1. Url.
	//  2. Responce type is "html" or "json" or "jsonp" or "text".
	//  3. The urltag is name for marking url and distinguish different urls in PageProcesser and Pipeline.
	//  4. The method is POST or GET.
	//  5. The postdata is body string sent to sever.
	//  6. The header is header for http request.
	//  7. Cookies
	req := request.NewRequest("http://bj.5i5j.com/exchange/124554654", "html", "", "GET", "", nil, nil, nil, nil)
	pageItems := sp.GetByRequest(req)

	url := pageItems.GetRequest().GetUrl()
	println("-----------------------------------spider.Get---------------------------------")
	println("url\t:\t" + url)
	for name, value := range pageItems.GetAll() {
		println(name + "\t:\t" + value)
	}

	// println("\n--------------------------------spider.GetAll---------------------------------")
	// urls := []string{
	// 	"http://baike.baidu.com/view/1628025.htm?fromtitle=http&fromid=243074&type=syn",
	// 	"http://baike.baidu.com/view/383720.htm?fromtitle=html&fromid=97049&type=syn",
	// }
	// var reqs []*request.Request
	// for _, url := range urls {
	// 	req := request.NewRequest(url, "html", "", "GET", "", nil, nil, nil, nil)
	// 	reqs = append(reqs, req)
	// }
	// pageItemsArr := sp.SetThreadnum(2).GetAllByRequest(reqs)
	// //pageItemsArr := sp.SetThreadnum(2).GetAll(urls, "html")
	// for _, item := range pageItemsArr {
	// 	url = item.GetRequest().GetUrl()
	// 	println("url\t:\t" + url)
	// 	fmt.Printf("item\t:\t%s\n", item.GetAll())
	// }
}
