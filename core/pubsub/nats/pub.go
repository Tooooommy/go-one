package nats

import (
	"context"
	"github.com/nats-io/nats.go"
)

type (
	Publisher struct {
		subject string
		reply   string
		conn    *Conn
	}
)

func (p *Publisher) PublishSync(ctx context.Context, data []byte) (interface{}, error) {
	return p.conn.Publish(ctx, &nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	})
}

func (p *Publisher) Publish(data []byte) error {
	return p.conn.PublishSync(&nats.Msg{
		Subject: p.subject,
		Reply:   p.reply,
		Data:    data,
	})
}
