package nats

import (
	"context"
	"github.com/Tooooommy/go-one/core/pubsub"
	"github.com/nats-io/nats.go"
	"time"
)

type (
	Conn struct {
		cfg  Config
		conn *nats.Conn
	}
)

func Connect(cfg Config) (*Conn, error) {
	var options []nats.Option
	if cfg.Name != "" {
		options = append(options, nats.Name(cfg.Name))
	}
	if cfg.Username != "" && cfg.Password != "" {
		options = append(options, nats.UserInfo(cfg.Username, cfg.Password))
	}

	if cfg.CertFile != "" && cfg.KeyFile != "" {
		options = append(options, nats.ClientCert(cfg.CertFile, cfg.KeyFile))
	}

	if cfg.Timeout <= 0 {
		cfg.Timeout = 10
	}
	options = append(options, nats.Timeout(time.Duration(cfg.Timeout)*time.Second))

	conn, err := nats.Connect(resolverAddr(cfg.Address), options...)
	if err != nil {
		return nil, err
	}
	return &Conn{cfg: cfg, conn: conn}, nil
}

func (c Conn) CONN() *nats.Conn {
	return c.conn
}

func (c Conn) CFG() Config {
	return c.cfg
}

func (c Conn) NewPublisher(subject string, reply string) pubsub.Publisher {
	p := &Publisher{conn: &c, subject: subject, reply: reply}
	return p
}

func (c Conn) Publish(ctx context.Context, msg *nats.Msg) (*nats.Msg, error) {
	return c.conn.RequestMsgWithContext(ctx, msg)
}

func (c Conn) PublishSync(msg *nats.Msg) error {
	return c.conn.PublishMsg(msg)
}

func (c Conn) NewSubscriber(subject, queue string) *Subscriber {
	return &Subscriber{conn: &c, subject: subject, queue: queue}
}

func (c Conn) Subscribe(subject, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return c.conn.QueueSubscribe(subject, queue, cb)
}
