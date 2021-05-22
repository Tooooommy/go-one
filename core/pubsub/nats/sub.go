package nats

import (
	"errors"
	"github.com/Tooooommy/go-one/core/pubsub"
	"github.com/nats-io/nats.go"
)

type (
	Subscriber struct {
		subject string
		queue   string
		conn    *Conn
		sub     *nats.Subscription
		err     error
	}
)

var (
	UnsubscribeErr = errors.New("pubsub subscriber unsubscribe")
)

// Subscribe
func (s *Subscriber) Subscribe(cb pubsub.MsgHandler) error {
	sub, err := s.conn.Subscribe(s.subject, s.queue, func(msg *nats.Msg) {
		defer msg.Ack()
		err := cb(msg.Data)
		if err != nil {
			s.err = err
		}
	})
	s.sub = sub
	return err
}

// Unsubscribe
func (s *Subscriber) Unsubscribe() error {
	if s.sub == nil {
		return UnsubscribeErr
	}
	return s.sub.Unsubscribe()
}
