package server

import "fmt"

var (
	DefaultName = "go-one"
	DefaultHost = "127.0.0.1"
	DefaultPort = 8081
)

type Conf struct {
	Name      string  `json:"name"`
	Host      string  `json:"host"`
	Port      int     `json:"port"`
	CertFile  string  `json:"cert_file"`
	KeyFile   string  `json:"key_file"`
	Secret    string  `json:"secret"`
	PreSecret string  `json:"pre_secret"`
	Limit     float64 `json:"limit"`
}

func (c *Conf) HaveCert() bool {
	return len(c.CertFile) > 0 && len(c.KeyFile) > 0
}

func (c *Conf) Address() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func DefaultConfig() Conf {
	return Conf{
		Name: DefaultName,
		Host: DefaultHost,
		Port: DefaultPort,
	}
}
