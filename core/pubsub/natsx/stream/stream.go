package stream

import (
	"github.com/Tooooommy/go-one/core/pubsub/natsx"
	"github.com/nats-io/nats.go"
)

type (
	Stream struct {
		stream nats.JetStreamContext
	}
)

// Connect
func Connect(conn *natsx.Conn) (*Stream, error) {
	stream, err := conn.CONN().JetStream()
	if err != nil {
		return nil, err
	}
	return &Stream{stream: stream}, nil
}

// Stream
func (s *Stream) Stream() nats.JetStreamContext {
	return s.stream
}
