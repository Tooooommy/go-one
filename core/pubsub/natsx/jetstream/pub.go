package jetstream

import (
	"github.com/nats-io/nats.go"
)

// publisher
type publisher struct {
	stream  *Stream
	subject string
	reply   string
	opts    []nats.PubOpt
}

// Publish
func (p *publisher) Publish(data []byte) (interface{}, error) {
	return p.stream.Publish(&nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	}, p.opts...)
}

// PublishSync
func (p *publisher) PublishSync(data []byte) (interface{}, error) {
	return p.stream.PublishSync(&nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	}, p.opts...)
}
