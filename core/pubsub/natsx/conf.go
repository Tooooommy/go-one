package natsx

import (
	"strings"
)

// Conf
type Conf struct {
	Name     string   `json:"name"`
	Address  []string `json:"address"`
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
func (cfg *Conf) Connect() Conn {
	return Connect(cfg)
}
