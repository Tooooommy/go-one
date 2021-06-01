package jaegerx

import (
	"github.com/Tooooommy/go-one/core/zapx"
)

type logger struct{}

func (*logger) Error(msg string) {
	zapx.Error().Msg(msg)
}

func (*logger) Infof(msg string, args ...interface{}) {
	zapx.Info().Any("args", args).Msg(msg)
}
