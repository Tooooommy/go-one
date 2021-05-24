package jetstream

import (
	"errors"
	"github.com/Tooooommy/go-one/core/pubsub"
	"github.com/nats-io/nats.go"
)

type subscriber struct {
	stream  *Stream
	subject string
	queue   string
	sub     *nats.Subscription
	opts    []nats.SubOpt
}

var (
	UnsubscribeErr = errors.New("pubsub subscriber unsubscribe")
)

func (s *subscriber) Subscribe(cb pubsub.MsgHandler) error {
	sub, err := s.stream.Subscribe(s.subject, s.queue, func(msg *nats.Msg) { cb(msg) }, s.opts...)
	s.sub = sub
	return err
}

func (s *subscriber) Unsubscribe() error {
	if s.sub == nil {
		return UnsubscribeErr
	}
	return s.sub.Unsubscribe()
}
