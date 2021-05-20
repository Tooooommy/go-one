package milvus

import "strings"

type Config struct {
	Address string `json:"address"`
}

func (cfg Config) DSN() string {
	return cfg.Address
}

func (cfg Config) NewClient() (*Client, error) {
	return NewClient(cfg)
}

func resolverAddr(address string) (host, port string) {
	ss := strings.Split(address, ":")
	if len(ss) >= 2 {
		host = ss[0]
		port = ss[1]
	} else if len(ss[0]) == 1 {
		port = "19530"
		if ss[0] == "" {
			host = "localhost"
		}
	}
	return
}
