package pubsub

import "context"

type (
	MsgHandler func([]byte) error

	Publisher interface {
		PublishSync(context.Context, []byte) (interface{}, error)
		Publish([]byte) error
	}

	Subscriber interface {
		Subscribe(MsgHandler) error
		Unsubscribe() error
		Error() error
	}
)
