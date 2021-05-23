package pubsub

type (
	//MsgHandler
	MsgHandler func(interface{})

	// Publisher
	Publisher interface {
		PublishSync([]byte) (interface{}, error)
		Publish([]byte) (interface{}, error)
	}

	// Subscriber
	Subscriber interface {
		Subscribe(MsgHandler) error
		Unsubscribe() error
	}
)
