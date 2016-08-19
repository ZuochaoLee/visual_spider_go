package pipeline

import (
	"github.com/Shopify/sarama"
	"strings"
	"time"
	"visual_spider_go/spider/core/common/com_interfaces"
	"visual_spider_go/spider/core/common/page_items"
)

type PipelineKafka struct {
	producer sarama.AsyncProducer
	db       string
}

// type ProducerMessage struct {
// 	Topic string
// 	Value []byte
// }

func NewPipelineKafka(host, db string) *PipelineKafka {
	brokerList := strings.Split(host, ",")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal       // Only wait for the leader to ack
	config.Producer.Compression = sarama.CompressionSnappy   // Compress messages
	config.Producer.Flush.Frequency = 500 * time.Millisecond // Flush batches every 500ms
	producer, _ := sarama.NewAsyncProducer(brokerList, config)
	return &PipelineKafka{producer: producer, db: db}
}
func (this *PipelineKafka) Process(items *page_items.PageItems, t com_interfaces.Task) {
	res := ""
	for k, v := range items.GetAll() {
		res = res + k + ":" + v + ","
	}
	var message sarama.ProducerMessage
	message.Topic = this.db
	message.Value = sarama.StringEncoder(res)
	this.producer.Input() <- &message
}
