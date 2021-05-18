package orm

import "github.com/meilisearch/meilisearch-go"

type MeiliConfig struct {
	Address string `json:"address"`
	ApiKey  string `json:"api_key"`
}

func (cfg MeiliConfig) DNS() meilisearch.Config {
	return meilisearch.Config{
		Host:   cfg.Address,
		APIKey: cfg.ApiKey,
	}
}

var (
	ldb meilisearch.ClientInterface
)

func InitMeili(cfg MeiliConfig) {
	ldb = meilisearch.NewClient(cfg.DNS())
}
