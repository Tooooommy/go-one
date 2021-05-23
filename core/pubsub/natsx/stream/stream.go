package stream

import (
	"github.com/Tooooommy/go-one/core/pubsub"
	"github.com/Tooooommy/go-one/core/pubsub/natsx"
	"github.com/nats-io/nats.go"
)

// Stream
type Stream struct {
	ctx nats.JetStreamContext
}

// Connect
func Connect(conn *natsx.Conn) (*Stream, error) {
	ctx, err := conn.CONN().JetStream()
	return &Stream{ctx: ctx}, err
}

// Context
func (s *Stream) Context() nats.JetStreamContext {
	return s.ctx
}

// NewPublisher
func (s *Stream) NewPublisher(subject string, reply string) pubsub.Publisher {
	return &publisher{
		stream:  s,
		subject: subject,
		reply:   reply,
	}
}

// Publish
func (s *Stream) Publish(msg *nats.Msg) (nats.PubAckFuture, error) {
	return s.ctx.PublishMsgAsync(msg)
}

// PublishSync
func (s *Stream) PublishSync(msg *nats.Msg) (*nats.PubAck, error) {
	return s.ctx.PublishMsg(msg)
}

// NewSubscriber
func (s *Stream) NewSubscriber(subject, queue string) pubsub.Subscriber {
	return &subscriber{stream: s, subject: subject, queue: queue}
}

// Subscribe
func (s *Stream) Subscribe(subject string, queue string, cb nats.MsgHandler) (*nats.Subscription, error) {
	return s.ctx.QueueSubscribe(subject, queue, cb)
}
