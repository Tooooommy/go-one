package orm

import (
	"database/sql"
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

var (
	db  *sql.DB
	gdb *gorm.DB
	xdb *sqlx.DB
)

func Ping(duration int) {
	defer func() {
		if result := recover(); result != nil {
			zapx.Error().Any("Recover Result", result).
				Msg("mysql ping function  recover")
		}
		Ping(duration)
	}()
	for {
		time.Sleep(time.Duration(duration) * time.Second)
		err := db.Ping()
		if err != nil {
			zapx.Error().Error(err).Msg("mysql database ping occurred error")
		}
	}
}

func initMysql(cfg Config) (err error) {
	db, err = sql.Open("mysql", cfg.MysqlDSN())
	if err != nil {
		return
	}

	db.SetConnMaxIdleTime(time.Duration(cfg.ConnMaxIdleTime) * time.Second)
	db.SetConnMaxLifetime(time.Duration(cfg.ConnMaxLifetime) * time.Second)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetMaxOpenConns(cfg.MaxOpenConns)
	Ping(cfg.PingDuration)
	return
}

func InitGorm(cfg Config) (err error) {
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

func InitSqlx(cfg Config) (err error) {
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
