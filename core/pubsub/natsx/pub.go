package natsx

import (
	"github.com/nats-io/nats.go"
	"time"
)

// publisher
type publisher struct {
	subject string
	reply   string
	timeout time.Duration
	conn    *Conn
}

// PublishSync
func (p *publisher) PublishSync(data []byte) (interface{}, error) {
	return p.conn.PublishSync(&nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	}, p.timeout)
}

// Publish
func (p *publisher) Publish(data []byte) (interface{}, error) {
	return nil, p.conn.Publish(&nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	})
}
