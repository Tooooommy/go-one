package mysql

import (
	"database/sql"
	"github.com/Tooooommy/go-one/core/syncx"
	"github.com/Tooooommy/go-one/core/zapx/gormx"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go.uber.org/zap/zapcore"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// client
type (
	Client interface {
		Conn() (*sql.DB, error)
	}

	client struct {
		cfg *Conf
	}
)

var (
	manager = syncx.NewManager()
)

// NewClient
func NewClient(cfg *Conf) Client {
	return &client{cfg: cfg}
}

func (c *client) getConn(dsn string) (*sql.DB, error) {
	val, exist := manager.Get("r-" + dsn)
	if exist {
		return val.(*sql.DB), nil
	}

	db, err := sql.Open("mysql", dsn)
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

	manager.Set("r-"+dsn, db)
	return db, nil
}

func (c *client) getXConn(dsn string) (*sqlx.DB, error) {
	val, exist := manager.Get("x-" + dsn)
	if exist {
		return val.(*sqlx.DB), nil
	}

	db, err := c.getConn(dsn)
	if err != nil {
		return nil, err
	}
	xdb := sqlx.NewDb(db, "mysql")
	manager.Set("x-"+dsn, xdb)
	return xdb, nil
}

func (c *client) getGConn(dsn string) (*gorm.DB, error) {
	val, exist := manager.Get("g-" + dsn)
	if exist {
		return val.(*gorm.DB), nil
	}

	db, err := c.getConn(dsn)
	if err != nil {
		return nil, err
	}
	log := logger.New(
		gormx.NewLogger(zapcore.InfoLevel),
		logger.Config{
			SlowThreshold:             0,
			Colorful:                  true,
			IgnoreRecordNotFoundError: true,
			LogLevel:                  logger.Info,
		},
	)
	nameing := schema.NamingStrategy{TablePrefix: c.cfg.TablePrefix}
	gdb, err := gorm.Open(
		mysql.New(mysql.Config{DriverName: "mysql", Conn: db}),
		&gorm.Config{NamingStrategy: nameing, Logger: log, ConnPool: db})
	manager.Set("g-"+dsn, gdb)
	return gdb, nil
}

// Conn
func (c *client) Conn() (*sql.DB, error) {
	return c.getConn(c.cfg.DSN())
}

// XConn
func (c *client) XConn() (*sqlx.DB, error) {
	return c.getXConn(c.cfg.DSN())
}

// GConn
func (c *client) GConn() (*gorm.DB, error) {
	return c.getGConn(c.cfg.DSN())
}
