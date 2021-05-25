package kafka

import (
	"github.com/Shopify/sarama"
)

type Conn struct {
	cfg  Config
	conn sarama.Client
}

func Connect(cfg Config) (*Conn, error) {
	conn, err := sarama.NewClient(cfg.Address, sarama.NewConfig())
	if err != nil {
		return nil, err
	}
	return &Conn{conn: conn, cfg: cfg}, nil
}

func (c *Conn) CONN() sarama.Client {
	return c.conn
}

func (c *Conn) CFG() Config {
	return c.cfg
}

func NewPublisher(conn *Conn, topic string) *publisher {
	return &publisher{
		topic:         topic,
		conn:          conn,
		syncProducer:  nil,
		asyncProducer: nil,
	}
}
