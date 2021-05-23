package stream

import (
	"github.com/nats-io/nats.go"
)

// publisher
type publisher struct {
	stream  *Stream
	subject string
	reply   string
}

// Publish
func (p *publisher) Publish(data []byte) (interface{}, error) {
	return p.stream.Publish(&nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	})
}

// PublishSync
func (p *publisher) PublishSync(data []byte) (interface{}, error) {
	return p.stream.PublishSync(&nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	})
}
