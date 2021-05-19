package meili

import "github.com/meilisearch/meilisearch-go"

func (p *Pool) Client() meilisearch.ClientInterface {
	return p.pool.Get().(meilisearch.ClientInterface)
}

func (p *Pool) Get() interface{} {
	return p.pool.Get()
}

func (p *Pool) Put(x interface{}) {
	p.pool.Put(x)
}
