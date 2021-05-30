package mysqlx

import "fmt"

type Config struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Database        string `json:"database"`
	Address         string `json:"address"`
	Charset         string `json:"charset"`
	Loc             string `json:"loc"`
	Timeout         int    `json:"timeout"`
	Prefix          string `json:"prefix"`
	MaxOpenConns    int    `json:"max_open_conns"`
	MaxIdleConns    int    `json:"max_idle_conns"`
	ConnMaxLifetime int    `json:"conn_max_lifetime"`
	ConnMaxIdleTime int    `json:"conn_max_idle_time"`
}

func (cfg *Config) DSN() string {
	username := cfg.Username
	password := cfg.Password
	database := cfg.Database
	address := cfg.Address
	charset := cfg.Charset
	loc := cfg.Loc
	timeout := cfg.Timeout

	if address == "" {
		address = "127.0.0.1:3306"
	}

	if charset == "" {
		charset = "utf8mb4"
	}

	if loc == "" {
		loc = "Local"
	}

	if timeout <= 0 {
		timeout = 10
	}

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s&timeout=%ds",
		username, password, address, database, charset, loc, 10)
}

func (cfg *Config) NewClient() (*Client, error) {
	return NewClient(cfg)
}
