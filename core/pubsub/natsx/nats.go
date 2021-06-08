package natsx

import (
	"github.com/Tooooommy/go-one/core/syncx"
	"github.com/nats-io/nats.go"
)

// conn
type (
	Conn interface {
		Conn() (*nats.Conn, error)
		Stream() (nats.JetStreamContext, error)
	}

	conn struct {
		cfg *Conf
	}
)

var (
	manager = syncx.NewManager()
)

// Connect
func Connect(cfg *Conf) Conn {
	return &conn{cfg: cfg}
}

func (c *conn) connect() (*nats.Conn, error) {
	addr := resolverAddr(c.cfg.Address)
	val, ok := manager.Get(addr)
	if ok {
		return val.(*nats.Conn), nil
	}
	conn, err := nats.Connect(addr, nats.Name(c.cfg.Name),
		nats.UserInfo(c.cfg.Username, c.cfg.Password))
	if err != nil {
		return nil, err
	}
	manager.Set(addr, conn)
	return conn, nil
}

// Conn
func (c *conn) Conn() (*nats.Conn, error) {
	return c.connect()
}

func (c *conn) Stream() (nats.JetStreamContext, error) {
	conn, err := c.connect()
	if err != nil {
		return nil, err
	}
	return conn.JetStream()
}
