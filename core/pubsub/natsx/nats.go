package natsx

import (
	"github.com/Tooooommy/go-one/core/pubsub"
	"github.com/nats-io/nats.go"
	"time"
)

// Conn
type Conn struct {
	cfg  Config
	conn *nats.Conn
}

// Connect
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

	c, err := nats.Connect(resolverAddr(cfg.Address), options...)
	if err != nil {
		return nil, err
	}
	conn := &Conn{cfg: cfg, conn: c}
	return conn, nil
}

// Close
func (c *Conn) Close() {
	c.conn.Close()
}

// CONN
func (c *Conn) CONN() *nats.Conn {
	return c.conn
}

// CFG
func (c Conn) CFG() Config {
	return c.cfg
}

// NewPublisher
func (c *Conn) NewPublisher(subject string, reply string, timeout time.Duration) pubsub.Publisher {
	return &publisher{conn: c, subject: subject, reply: reply, timeout: timeout}
}

// Publish
func (c *Conn) PublishSync(msg *nats.Msg, timeout time.Duration) (*nats.Msg, error) {
	return c.conn.RequestMsg(msg, timeout)
}

// PublishSync
func (c *Conn) Publish(msg *nats.Msg) error {
	return c.conn.PublishMsg(msg)
}

// NewSubscriber
func (c *Conn) NewSubscriber(subject, queue string) pubsub.Subscriber {
	return &subscriber{conn: c, subject: subject, queue: queue}
}

// Subscribe
func (c *Conn) Subscribe(subject, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return c.conn.QueueSubscribe(subject, queue, cb)
}
