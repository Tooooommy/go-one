package natsx

import (
	"context"
	"github.com/nats-io/nats.go"
)

type (
	publisher struct {
		subject string
		reply   string
		conn    *Conn
	}
)

func (p *publisher) PublishSync(ctx context.Context, data []byte) (interface{}, error) {
	return p.conn.PublishSync(ctx, &nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	})
}

func (p *publisher) Publish(data []byte) error {
	return p.conn.Publish(&nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	})
}
