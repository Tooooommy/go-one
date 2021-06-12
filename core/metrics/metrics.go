package metrics

import kprovider "github.com/go-kit/kit/metrics/provider"

type Metrics interface {
	kprovider.Provider
	Namespace() string
	Subsystem() string
}
