package kafka

import (
	"github.com/Shopify/sarama"
)

type publisher struct {
	topic         string
	conn          *Conn
	syncProducer  sarama.SyncProducer
	asyncProducer sarama.AsyncProducer
}

// 异步
func (p *publisher) Publish([]byte) (interface{}, error) {
	return nil, nil
}
