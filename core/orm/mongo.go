package orm

import (
	"context"
	"fmt"
	"github.com/Tooooommy/go-one/core/zapx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
	"time"
)

type MongoConfig struct {
	Username string            `json:"username"`
	Password string            `json:"password"`
	Address  []string          `json:"address"`
	Database string            `json:"database"`
	Options  map[string]string `json:"options"`
}

func (cfg MongoConfig) DSN() string {
	username := cfg.Username
	password := cfg.Password
	address := strings.Join(cfg.Address, ",")
	database := "test"

	if len(cfg.Address) <= 0 {
		address = "127.0.0.1:27017"
	}
	if cfg.Database != "" {
		database = cfg.Database
	}
	var opts []string
	for k, v := range cfg.Options {
		opts = append(opts, k+"="+v)
	}
	opt := strings.Join(opts, "&")
	return fmt.Sprintf("mongdb://%s:%s@%s/%s?%s", username, password, address, database, opt)
}

var (
	mdb *mongo.Client
)

func PingMongo(duration int) {
	defer func() {
		if result := recover(); result != nil {
			zapx.Error().Any("Recover Result", result).
				Msg("mongo ping function recover")
		}
		PingMongo(duration)
	}()
	for {
		time.Sleep(time.Duration(duration) * time.Second)
		err := mdb.Ping(context.Background(), readpref.Primary())
		if err != nil {
			zapx.Error().Error(err).Msg("mongo database ping occurred error")
		}
	}
}

func initMongo(cfg MongoConfig) (err error) {
	uri := options.Client().ApplyURI(cfg.DSN())
	mdb, err = mongo.Connect(context.Background(), uri)
	if err != nil {
		return
	}
	return
}

func GetMongoAuto() *mongo.Client {
	return mdb
}
