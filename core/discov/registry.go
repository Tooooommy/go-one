package discov

type (

	// Registry
	Registry interface {
		Register(Service) error
		Deregister(Service) error
	}

	// Service
	Service struct {
		Key       string `json:"key"`
		Val       string `json:"val"`
		Heartbeat int64  `json:"heartbeat"`
		TTL       int64  `json:"ttl"`
	}
)

func (s Service) Register(r Registry) (err error) {
	if r == nil {
		return nil
	}
	defer func() {
		err = r.Deregister(s)
	}()
	err = r.Register(s)
	return
}
