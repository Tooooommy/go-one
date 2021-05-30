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
	MaxOpenConns    int    `json:"max_open_conns"`
	MaxIdleConns    int    `json:"max_idle_conns"`
	ConnMaxLifetime int    `json:"conn_max_lifetime"`
	ConnMaxIdleTime int    `json:"conn_max_idle_time"`
}

func DefaultConfig() *Config {
	return &Config{
		Username: "root",
		Password: "root",
		Database: "master",
		Address:  "127.0.0.1:3306",
		Charset:  "utf8mb4",
		Loc:      "Local",
		Timeout:  10,
	}
}

func (cfg *Config) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s&timeout=%ds",
		cfg.Username, cfg.Password, cfg.Address, cfg.Database, cfg.Charset, cfg.Loc, cfg.Timeout)
}
