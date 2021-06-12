package mysql

import (
	"github.com/Tooooommy/go-one/core/syncx"
	"github.com/Tooooommy/go-one/core/zapx/gormx"
	_ "github.com/go-sql-driver/mysql"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// client
type (
	Client interface {
		Conn() (*gorm.DB, error)
	}

	client struct {
		cfg *Conf
	}
)

var (
	manager = syncx.NewManager()
)

// NewClient new client instance contains conf
func NewClient(cfg *Conf) Client {
	return &client{cfg: cfg}
}

// getConn get client gorm db from dsn:conf
func (c *client) getConn(dsn string) (*gorm.DB, error) {
	val, ok := manager.Get(dsn)
	if ok {
		return val.(*gorm.DB), nil
	}

	gdb, err := gorm.Open(mysql.Open(dsn), getGormConfig(c.cfg.TablePrefix))
	if err != nil {
		return nil, err
	}

	db, err := gdb.DB()
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if c.cfg.ConnMaxIdleTime > 0 {
		db.SetConnMaxIdleTime(c.cfg.ConnMaxIdleTime)
	}

	if c.cfg.ConnMaxLifetime > 0 {
		db.SetConnMaxLifetime(c.cfg.ConnMaxLifetime)
	}

	if c.cfg.MaxIdleConns > 0 {
		db.SetMaxIdleConns(c.cfg.MaxIdleConns)
	}

	if c.cfg.MaxOpenConns > 0 {
		db.SetMaxOpenConns(c.cfg.MaxOpenConns)
	}

	manager.Set(dsn, gdb)
	return gdb, nil
}

// Conn
func (c *client) Conn() (*gorm.DB, error) {
	return c.getConn(c.cfg.DSN())
}

func getGormConfig(tablePrefix string) *gorm.Config {
	return &gorm.Config{NamingStrategy: schema.NamingStrategy{TablePrefix: tablePrefix},
		Logger: logger.New(
			gormx.NewLogger(zapcore.InfoLevel),
			logger.Config{
				SlowThreshold:             0,
				Colorful:                  true,
				IgnoreRecordNotFoundError: true,
				LogLevel:                  logger.Info,
			},
		),
	}
}
