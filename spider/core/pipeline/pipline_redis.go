package pipeline

import (
	redis "github.com/alphazero/Go-Redis"
	"math/rand"
	"strconv"
	"time"
	"visual_spider_go/spider/core/common/com_interfaces"
	"visual_spider_go/spider/core/common/page_items"
)

type PipelineRedis struct {
	client redis.Client
}

func NewPipelineRedis(host string, port int, db int, passwd string) *PipelineRedis {
	spec := redis.DefaultSpec().Host(host).Port(port).Db(db).Password(passwd)
	client, _ := redis.NewSynchClientWithSpec(spec)
	return &PipelineRedis{client: client}
}
func (this *PipelineRedis) Process(items *page_items.PageItems, t com_interfaces.Task) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ts := r.Intn(100)
	ta := time.Now().Unix()
	tt := strconv.FormatInt(ta, 10) + strconv.Itoa(ts)
	for k, v := range items.GetAll() {
		this.client.Hset(tt, k, []byte(v))
	}
}
