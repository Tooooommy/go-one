package jetstream

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
func (s *Stream) NewPublisher(subject string, reply string, opts ...nats.PubOpt) pubsub.Publisher {
	return &publisher{
		stream:  s,
		subject: subject,
		reply:   reply,
		opts:    opts,
	}
}

// Publish
func (s *Stream) Publish(msg *nats.Msg, opts ...nats.PubOpt) (nats.PubAckFuture, error) {
	return s.ctx.PublishMsgAsync(msg, opts...)
}

// PublishSync
func (s *Stream) PublishSync(msg *nats.Msg, opts ...nats.PubOpt) (*nats.PubAck, error) {
	return s.ctx.PublishMsg(msg, opts...)
}

// NewSubscriber
func (s *Stream) NewSubscriber(subject, queue string, opts ...nats.SubOpt) pubsub.Subscriber {
	return &subscriber{stream: s, subject: subject, queue: queue, opts: opts}
}

// Subscribe
func (s *Stream) Subscribe(subject string, queue string, cb nats.MsgHandler, opts ...nats.SubOpt) (*nats.Subscription, error) {
	return s.ctx.QueueSubscribe(subject, queue, cb, opts...)
}
