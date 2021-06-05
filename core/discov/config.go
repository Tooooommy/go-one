package discov

import (
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

type Config struct {
	Name          string   `json:"name"`      // key
	Hosts         []string `json:"endpoints"` // val
	Username      string   `json:"username"`
	Password      string   `json:"password"`
	CertFile      string   `json:"cert_file"`
	KeyFile       string   `json:"key_file"`
	DialTimeout   int64    `json:"dial_timeout"`
	DialKeepAlive int64    `json:"dial_keep_alive"`
	Heartbeat     int64    `json:"heartbeat"`
	Ttl           int64    `json:"ttl"`
}

func (c *Config) ClientOptions() etcdv3.ClientOptions {
	return etcdv3.ClientOptions{
		Cert:          c.CertFile,
		Key:           c.KeyFile,
		DialTimeout:   time.Duration(c.DialTimeout) * time.Second,
		DialKeepAlive: time.Duration(c.DialKeepAlive) * time.Second,
		Username:      c.Username,
		Password:      c.Password,
	}
}

func (c Config) HaveEtcd() bool {
	return len(c.Hosts) > 0 && len(c.Name) > 0
}
