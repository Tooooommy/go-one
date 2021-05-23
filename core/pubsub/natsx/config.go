package natsx

import (
	"strings"
)

// Config
type Config struct {
	Name     string   `json:"name"`
	Address  []string `json:"address"`
	CertFile string   `json:"cert_file"`
	KeyFile  string   `json:"key_file"`
	Timeout  int64    `json:"timeout"`
	Username string   `json:"username"`
	Password string   `json:"password"`
}

// resolverAddr
func resolverAddr(address []string) string {
	for index := range address {
		if !strings.HasPrefix(address[index], "natsx://") {
			address[index] = "natsx://" + address[index]
		}
	}
	return strings.Join(address, ",")
}

// Connect
func (cfg Config) Connect() (*Conn, error) {
	return Connect(cfg)
}
