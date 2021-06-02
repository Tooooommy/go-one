package etcdx

//import (
//	"github.com/go-kit/kit/sd/etcdv3"
//	"time"
//)
//
//type Config struct {
//	Hosts         []string `json:"hosts"`
//	CertFile      string   `json:"cert_file"`
//	KeyFile       string   `json:"key_file"`
//	DialTimeout   int64    `json:"dial_timeout"`
//	DialKeepAlive int64    `json:"dial_keep_alive"`
//	Username      string   `json:"username"`
//	Password      string   `json:"password"`
//}
//
//func (c *Config) ClientOptions() etcdv3.ClientOptions {
//	return etcdv3.ClientOptions{
//		Cert:          c.CertFile,
//		Key:           c.KeyFile,
//		DialTimeout:   time.Duration(c.DialTimeout),
//		DialKeepAlive: time.Duration(c.DialKeepAlive),
//		Username:      c.Username,
//		Password:      c.Password,
//	}
//}
//
//func (c Config) HaveEtcd() bool {
//	return len(c.Hosts) > 0
//}
