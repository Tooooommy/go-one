package logx

import (
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
	cfg Config
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
		cfg: Config{},
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

func SetStdMode(cfg StdModeConfig) Option {
	return func(zapx *Zapx) {
		l := zap.NewAtomicLevelAt(zapcore.Level(cfg.Level))
		w := zapcore.AddSync(os.Stdout)
		c := zap.NewProductionEncoderConfig()
		c.EncodeTime = zapcore.ISO8601TimeEncoder
		core := zapcore.NewCore(zapcore.NewJSONEncoder(c), w, l)
		_zapx.log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
		_zapx.cfg = cfg.Config
	}
}

func SetLogMode(cfg LogModeConfig) Option {
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
		_zapx.cfg = cfg.Config
	}
}

func SetElkMode(cfg ElkModeConfig) Option {
	return func(zapx *Zapx) {
	}
}

func KitL() log.Logger {
	return kitlog.NewZapSugarLogger(_zapx.log, zapcore.Level(_zapx.cfg.Level))
}

func L() *zap.Logger {
	return _zapx.log
}

func S() *zap.SugaredLogger {
	return _zapx.sug
}

func Debug() ZapxLogger {
	return newZapx(zapcore.DebugLevel)
}

func Info() ZapxLogger {
	return newZapx(zapcore.InfoLevel)
}

func Warn() ZapxLogger {
	return newZapx(zapcore.WarnLevel)
}

func Error() ZapxLogger {
	return newZapx(zapcore.ErrorLevel)
}

func Dpanic() ZapxLogger {
	return newZapx(zapcore.DPanicLevel)
}

func Panic() ZapxLogger {
	return newZapx(zapcore.PanicLevel)
}

func Fatal() ZapxLogger {
	return newZapx(zapcore.FatalLevel)
}
