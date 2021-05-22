package nats

import (
	"strings"
)

type Config struct {
	Name           string   `json:"name"`
	Address        []string `json:"address"`
	CertFile       string   `json:"cert_file"`
	KeyFile        string   `json:"key_file"`
	Timeout        int64    `json:"timeout"`
	Username       string   `json:"username"`
	Password       string   `json:"password"`
	PublishTimeout int64    `json:"publish_timeout"`
}

func resolverAddr(address []string) string {
	for index := range address {
		if !strings.HasPrefix(address[index], "nats://") {
			address[index] = "nats://" + address[index]
		}
	}
	return strings.Join(address, ",")
}

func (cfg Config) Connect() {
	Connect(cfg)
}
