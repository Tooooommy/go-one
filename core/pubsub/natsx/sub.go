package natsx

import (
	"errors"
	"github.com/Tooooommy/go-one/core/pubsub"
	"github.com/nats-io/nats.go"
)

// subscriber
type subscriber struct {
	subject string
	queue   string
	conn    *Conn
	sub     *nats.Subscription
}

// UnsubscribeErr
var UnsubscribeErr = errors.New("pubsub subscriber unsubscribe")

// Subscribe
func (s *subscriber) Subscribe(cb pubsub.MsgHandler) error {
	sub, err := s.conn.Subscribe(s.subject, s.queue, func(msg *nats.Msg) { cb(msg) })
	s.sub = sub
	return err
}

// Unsubscribe
func (s *subscriber) Unsubscribe() error {
	if s.sub == nil {
		return UnsubscribeErr
	}
	return s.sub.Unsubscribe()
}
