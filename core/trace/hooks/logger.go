package hooks

import (
	"context"
	"fmt"
	"github.com/Tooooommy/go-one/core/zapx"
)

type logger struct{}

func NewLogger() *logger {
	return &logger{}
}

func (*logger) Error(msg string) {
	zapx.Error(context.Background()).Msg(msg)
}

func (*logger) Infof(msg string, args ...interface{}) {
	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args)
	}
	zapx.Info(context.Background()).Msg(msg)
}
