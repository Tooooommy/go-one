package logx

import (
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestZapx(t *testing.T) {
	Logger().Core().Enabled(zapcore.DebugLevel)
	Debug().String("hello", "world").Msg("say")
}
