package gormx

import (
	"github.com/Tooooommy/go-one/core/zapx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	sugar *zap.SugaredLogger
	level zapcore.Level
}

func NewLogger(level zapcore.Level) *Logger {
	return &Logger{
		sugar: zapx.S(),
		level: level,
	}
}

func (l *Logger) Printf(template string, args ...interface{}) {
	switch l.level {
	case zapcore.DebugLevel:
		l.sugar.Debug(template, args)
	case zapcore.InfoLevel:
		l.sugar.Infof(template, args...)
	case zapcore.WarnLevel:
		l.sugar.Warnf(template, args...)
	case zapcore.FatalLevel:
		l.sugar.Fatalf(template, args...)
	case zapcore.PanicLevel:
		l.sugar.Panicf(template, args...)
	case zapcore.DPanicLevel:
		l.sugar.DPanicf(template, args...)
	case zapcore.ErrorLevel:
		l.sugar.Errorf(template, args...)
	default:
		l.sugar.Infof(template, args...)
	}
}
