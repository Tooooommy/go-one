package jaegerx

import "github.com/Tooooommy/go-one/core/zapx"

type logger struct{}

func (*logger) Error(msg string) {
	zapx.S().Error(msg)
}

func (*logger) Infof(msg string, args ...interface{}) {
	zapx.S().Infof(msg, args...)
}
