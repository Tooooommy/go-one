package mysql

import (
	"fmt"
	"time"
)

type Conf struct {
	Username        string        `json:"username"`
	Password        string        `json:"password"`
	Database        string        `json:"database"`
	Address         string        `json:"address"`
	Charset         string        `json:"charset"`
	Loc             string        `json:"loc"`
	Timeout         time.Duration `json:"timeout"`
	MaxOpenConns    int           `json:"max_open_conns"`
	MaxIdleConns    int           `json:"max_idle_conns"`
	ConnMaxLifetime time.Duration `json:"conn_max_lifetime"`
	ConnMaxIdleTime time.Duration `json:"conn_max_idle_time"`
	TablePrefix     string        `json:"table_prefix"`
}

func DefaultConf() *Conf {
	return &Conf{
		Username:        "root",
		Password:        "root",
		Database:        "master",
		Address:         "127.0.0.1:3306",
		Charset:         "utf8mb4",
		Loc:             "Local",
		Timeout:         10,
		MaxOpenConns:    64,
		MaxIdleConns:    64,
		ConnMaxLifetime: 60,
		ConnMaxIdleTime: 60,
	}
}

func (cfg *Conf) DSN() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=True&loc=%s&timeout=%ds",
		cfg.Username, cfg.Password, cfg.Address, cfg.Database, cfg.Charset, cfg.Loc, cfg.Timeout)
}

func (cfg *Conf) NewClient() Client {
	return NewClient(cfg)
}
