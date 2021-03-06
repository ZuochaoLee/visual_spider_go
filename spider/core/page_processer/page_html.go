package page_processer

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"visual_spider_go/spider/core/common/page"
)

type PageProcesserHtml struct {
	conf map[string]string
	page map[string]string
	rule map[string]string
	fun  map[string]string
	num  map[string]string
}

func NewPageProcesserHtml(conf, page, rule, fun, num map[string]string) *PageProcesserHtml {

	return &PageProcesserHtml{conf: conf, page: page, rule: rule, fun: fun, num: num}
}

func (this *PageProcesserHtml) Process(p *page.Page) {
	if !p.IsSucc() {
		println(p.Errormsg())
		return
	}

	result := map[string]string{}
	for k, _ := range this.rule {
		result[k] = ""
	}

	query := p.GetHtmlParser()
	var urls []string
	query.Find(this.page["rule"]).Each(func(i int, s *goquery.Selection) {
		href := ""
		if this.page["fun"] == "text" {
			href = s.Text()
		} else {
			href, _ = s.Attr(this.page["fun"])
		}

		urls = append(urls, this.page["pre"]+href)
	})
	p.AddMyTargetRequests(urls, this.conf["texttype"], "", this.conf["resqType"], this.conf["postdata"], this.conf["proxy"], this.conf["heardefile"], this.conf["cookie"])
	for k, v := range this.rule {
		if this.num[k] == "ALL" {
			var items []string
			query.Find(v).Each(func(i int, s *goquery.Selection) {
				item := ""
				if this.fun[k] == "text" {
					item = s.Text()
				} else {
					item, _ = s.Attr(this.fun[k])
				}
				items = append(items, item)
			})
			result[k] = strings.Join(items, "|")
		} else {
			if this.fun[k] == "text" {
				result[k] = query.Find(v).Text()
			} else {
				result[k], _ = query.Find(v).Attr(this.fun[k])
			}
			result[k] = strings.Trim(result[k], " \t\n")
		}

		if result[k] == "" {
			p.SetSkip(true)
		}
		p.AddField(k, result[k])

	}
	for k, v := range p.GetPageItems().GetAll() {
		println(k, v)
	}

}

func (this *PageProcesserHtml) Finish() {

	fmt.Printf("TODO:before end spider \r\n")
}
