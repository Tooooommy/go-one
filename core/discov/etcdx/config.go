package etcdx

import (
	"github.com/go-kit/kit/sd/etcdv3"
	"time"
)

type Config struct {
	Hosts         []string `json:"hosts"`
	Cert          string   `json:"cert"`
	Key           string   `json:"key"`
	CACert        string   `json:"ca_cert"`
	DialTimeout   int64    `json:"dial_timeout"`
	DialKeepAlive int64    `json:"dial_keep_alive"`
	Username      string   `json:"username"`
	Password      string   `json:"password"`
}

func (c Config) GetEtcdClientOptions() etcdv3.ClientOptions {
	return etcdv3.ClientOptions{
		Cert:          c.Cert,
		Key:           c.Key,
		CACert:        c.CACert,
		DialTimeout:   time.Duration(c.DialTimeout),
		DialKeepAlive: time.Duration(c.DialKeepAlive),
		Username:      c.Username,
		Password:      c.Password,
	}
}

func (c Config) HaveEtcd() bool {
	return len(c.Hosts) > 0
}
