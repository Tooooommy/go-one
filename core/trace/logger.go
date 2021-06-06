package trace

import (
	"fmt"
	"github.com/Tooooommy/go-one/core/zapx"
)

type logger struct{}

func (*logger) Error(msg string) {
	zapx.Error().Msg(msg)
}

func (*logger) Infof(msg string, args ...interface{}) {
	msg = fmt.Sprintf(msg, args)
	zapx.Info().Msg(msg)
}
