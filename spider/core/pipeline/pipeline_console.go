package pipeline

import (
	"../common/com_interfaces"
	"../common/page_items"
)

type PipelineConsole struct {
}

func NewPipelineConsole() *PipelineConsole {
	return &PipelineConsole{}
}

func (this *PipelineConsole) Process(items *page_items.PageItems, t com_interfaces.Task) {
	println("----------------------------------------------------------------------------------------------")
	println("Crawled url :\t" + items.GetRequest().GetUrl() + "\n")
	println("Crawled result : ")
	for key, value := range items.GetAll() {
		println(key + "\t:\t" + value)
	}
}
