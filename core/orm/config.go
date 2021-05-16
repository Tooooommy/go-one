package orm

import "fmt"

type Config struct {
	Username        string `json:"username"`
	Password        string `json:"password"`
	Database        string `json:"database"`
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Charset         string `json:"charset"`
	Loc             string `json:"loc"`
	Timeout         int    `json:"timeout"`
	Prefix          string `json:"prefix"`
	MaxOpenConns    int    `json:"max_open_conns"`
	MaxIdleConns    int    `json:"max_idle_conns"`
	ConnMaxLifetime int    `json:"conn_max_lifetime"`
	ConnMaxIdleTime int    `json:"conn_max_idle_time"`
	PingDuration    int    `json:"ping_duration"`
}

func (cfg Config) MysqlDSN() string {
	username := cfg.Username
	password := cfg.Password
	database := cfg.Database

	host := cfg.Host
	if cfg.Host == "" {
		host = "127.0.0.1"
	}

	port := cfg.Port
	if cfg.Port <= 0 {
		port = 3306
	}

	charset := cfg.Charset
	if charset == "" {
		charset = "utf8mb4"
	}

	loc := cfg.Loc
	if loc == "" {
		loc = "Local"
	}
	timeout := cfg.Timeout
	if timeout <= 0 {
		timeout = 10
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=%s&timeout=%ds",
		username, password, host, port, database, charset, loc, 10)
}
