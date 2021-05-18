package orm

import (
	"database/sql"
	"fmt"
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/Tooooommy/go-one/core/zapx/gormx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

type MySqlConfig struct {
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
	PingDuration    int    `json:"ping_duration"`
}

func (cfg MySqlConfig) DSN() string {
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

var (
	db  *sql.DB
	gdb *gorm.DB
	xdb *sqlx.DB
)

func PingMysql(duration int) {
	defer func() {
		if result := recover(); result != nil {
			zapx.Error().Any("Recover Result", result).
				Msg("mysql ping function recover")
		}
		PingMysql(duration)
	}()
	for {
		time.Sleep(time.Duration(duration) * time.Second)
		err := db.Ping()
		if err != nil {
			db = nil
			zapx.Error().Error(err).Msg("mysql database ping occurred error")
		}
	}
}

func initMysql(cfg MySqlConfig) (err error) {
	db, err = sql.Open("mysql", cfg.DSN())
	if err != nil {
		return
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	PingMysql(cfg.PingDuration)
	return
}

func InitGorm(cfg MySqlConfig) (err error) {
	if db == nil {
		err = initMysql(cfg)
		if err != nil {
			return
		}
	}
	log := logger.New(gormx.NewLogger(zapcore.InfoLevel), logger.Config{
		SlowThreshold:             0,
		Colorful:                  true,
		IgnoreRecordNotFoundError: true,
		LogLevel:                  logger.Info,
	})
	gdb, err = gorm.Open(mysql.New(mysql.Config{
		DriverName: "mysql",
		Conn:       db,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: cfg.Prefix,
		},
		Logger:   log,
		ConnPool: db,
	})
	return
}

func GetGormAuto() *gorm.DB {
	return gdb
}

func InitSqlx(cfg MySqlConfig) (err error) {
	if db == nil {
		err = initMysql(cfg)
		if err != nil {
			return
		}
	}
	xdb = sqlx.NewDb(db, "mysql")
	return nil
}

func GetSqlxAuto() *sqlx.DB {
	return xdb
}
