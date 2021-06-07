package zapx

import (
	"context"
	"github.com/Tooooommy/go-one/tools"
	"github.com/go-kit/kit/log"
	kitlog "github.com/go-kit/kit/log/zap"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func init() {
	NewZapx()
}

type Zapx struct {
	log *zap.Logger
	sug *zap.SugaredLogger
	cfg *Config
}

type Option func(*Zapx)

var _zapx *Zapx

func NewZapx(options ...Option) {
	l, err := zap.NewDevelopment(
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	)
	if err != nil {
		panic(err)
	}
	_zapx = &Zapx{
		log: l,
		sug: l.Sugar(),
		cfg: &Config{},
	}
	for _, opt := range options {
		opt(_zapx)
	}
	if _zapx.cfg.Name != "" {
		_zapx.log = _zapx.log.With(zap.String("name", _zapx.cfg.Name))
	}
	_zapx.sug = _zapx.log.Sugar()
	zap.ReplaceGlobals(_zapx.log)()
}

func SetStdMode(cfg *StdModeConfig) Option {
	return func(zapx *Zapx) {
		l := zap.NewAtomicLevelAt(zapcore.Level(cfg.Level))
		w := zapcore.AddSync(os.Stdout)
		c := zap.NewProductionEncoderConfig()
		c.EncodeTime = zapcore.ISO8601TimeEncoder
		core := zapcore.NewCore(zapcore.NewJSONEncoder(c), w, l)
		_zapx.log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		_zapx.cfg = &cfg.Config
	}
}

func SetLogMode(cfg *LogModeConfig) Option {
	return func(zapx *Zapx) {
		l := zap.NewAtomicLevelAt(zapcore.Level(cfg.Level))
		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.MaxSize,
			MaxAge:     cfg.MaxAge,
			MaxBackups: cfg.MaxBackups,
			LocalTime:  cfg.LocalTime,
			Compress:   cfg.Compress,
		})
		c := zap.NewProductionEncoderConfig()
		c.EncodeTime = zapcore.ISO8601TimeEncoder
		core := zapcore.NewCore(zapcore.NewJSONEncoder(c), w, l)
		_zapx.log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		_zapx.cfg = &cfg.Config
	}
}

func SetElkMode(cfg *ElkModeConfig) Option {
	return func(zapx *Zapx) {
	}
}

func KitL() log.Logger {
	return kitlog.NewZapSugarLogger(_zapx.log, zapcore.Level(_zapx.cfg.Level))
}

func L(ctx context.Context) *zap.Logger {
	key, name, err := tools.ExtractTraceKeyFromCtx(ctx)
	if err == nil {
		_zapx.log = _zapx.log.With(
			zap.String("trace_key", key),
			zap.String("span_name", name))
	} else if err != tools.ErrNotExistTraceSpan {
		_zapx.log = _zapx.log.With(
			zap.String("trace_err", err.Error()))
	}
	return _zapx.log
}

func S(ctx context.Context) *zap.SugaredLogger {
	key, name, err := tools.ExtractTraceKeyFromCtx(ctx)
	if err == nil {
		_zapx.sug = _zapx.sug.
			With("trace_key", key).
			With("span_name", name)
	} else if err != tools.ErrNotExistTraceSpan {
		_zapx.sug = _zapx.sug.With("trace_err", err.Error())
	}
	return _zapx.sug
}

func Debug(ctx context.Context) ZapxLogger {
	return newZapx(ctx, zapcore.DebugLevel)
}

func Info(ctx context.Context) ZapxLogger {
	return newZapx(ctx, zapcore.InfoLevel)
}

func Warn(ctx context.Context) ZapxLogger {
	return newZapx(ctx, zapcore.WarnLevel)
}

func Error(ctx context.Context) ZapxLogger {
	return newZapx(ctx, zapcore.ErrorLevel)
}

func Dpanic(ctx context.Context) ZapxLogger {
	return newZapx(ctx, zapcore.DPanicLevel)
}

func Panic(ctx context.Context) ZapxLogger {
	return newZapx(ctx, zapcore.PanicLevel)
}

func Fatal(ctx context.Context) ZapxLogger {
	return newZapx(ctx, zapcore.FatalLevel)
}
