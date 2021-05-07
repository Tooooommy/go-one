package server

var (
	DefaultName = "go-one"
	DefaultHost = "127.0.0.1"
	DefaultPort = 8081
)

type Config struct {
	Name     string `json:"name"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	CertFile string `json:"cert_file"`
	KeyFile  string `json:"key_file"`
}

func DefaultConfig() Config {
	return Config{
		Name: DefaultName,
		Host: DefaultHost,
		Port: DefaultPort,
	}
}
