package meili

import (
	"github.com/meilisearch/meilisearch-go"
	"sync"
)

// Pool
type Pool struct {
	Address string
	ApiKey  string
	pool    sync.Pool
}

// newMeili
func newMeili(address, apikey string) meilisearch.ClientInterface {
	return meilisearch.NewClient(meilisearch.Config{
		Host:   address,
		APIKey: apikey,
	})
}

// NewPool
func NewPool(address, apikey string) *Pool {
	return &Pool{
		Address: address,
		ApiKey:  apikey,
		pool: sync.Pool{
			New: func() interface{} {
				return newMeili(address, apikey)
			},
		},
	}
}
