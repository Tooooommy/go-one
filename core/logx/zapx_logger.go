package logx

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

type ZapxLogger interface {
	Skip() ZapxLogger
	Stack(key string) ZapxLogger
	StackSkip(key string, skip int) ZapxLogger
	Binary(key string, val []byte) ZapxLogger
	Array(key string, val zapcore.ArrayMarshaler) ZapxLogger
	Bool(key string, val bool) ZapxLogger
	Boolp(key string, val *bool) ZapxLogger
	Bools(key string, val []bool) ZapxLogger
	ByteString(key string, val []byte) ZapxLogger
	ByteStrings(key string, val [][]byte) ZapxLogger
	Complex128(key string, val complex128) ZapxLogger
	Complex128p(key string, val *complex128) ZapxLogger
	Complex128s(key string, val []complex128) ZapxLogger
	Complex64(key string, val complex64) ZapxLogger
	Complex64p(key string, val *complex64) ZapxLogger
	Complex64s(key string, val []complex64) ZapxLogger
	Float64(key string, val float64) ZapxLogger
	Float64p(key string, val *float64) ZapxLogger
	Float64s(key string, val []float64) ZapxLogger
	Float32(key string, val float32) ZapxLogger
	Float32p(key string, val *float32) ZapxLogger
	Float32s(key string, val []float32) ZapxLogger
	Int(key string, val int) ZapxLogger
	Intp(key string, val *int) ZapxLogger
	Ints(key string, val []int) ZapxLogger
	Int64(key string, val int64) ZapxLogger
	Int64p(key string, val *int64) ZapxLogger
	Int64s(key string, val []int64) ZapxLogger
	Int32(key string, val int32) ZapxLogger
	Int32p(key string, val *int32) ZapxLogger
	Int32s(key string, val []int32) ZapxLogger
	Int16(key string, val int16) ZapxLogger
	Int16p(key string, val *int16) ZapxLogger
	Int16s(key string, val []int16) ZapxLogger
	Int8(key string, val int8) ZapxLogger
	Int8p(key string, val *int8) ZapxLogger
	Int8s(key string, val []int8) ZapxLogger
	Unit(key string, val uint) ZapxLogger
	Unitp(key string, val *uint) ZapxLogger
	Units(key string, val []uint) ZapxLogger
	Uint64(key string, val uint64) ZapxLogger
	Uint64p(key string, val *uint64) ZapxLogger
	Uint64s(key string, val []uint64) ZapxLogger
	Uint32(key string, val uint32) ZapxLogger
	Uint32p(key string, val *uint32) ZapxLogger
	Uint32s(key string, val []uint32) ZapxLogger
	Uint16(key string, val uint16) ZapxLogger
	Uint16p(key string, val *uint16) ZapxLogger
	Uint16s(key string, val []uint16) ZapxLogger
	Uint8(key string, val uint8) ZapxLogger
	Uint8p(key string, val *uint8) ZapxLogger
	Uint8s(key string, val []uint8) ZapxLogger
	Uintptr(key string, val uintptr) ZapxLogger
	Uintptrp(key string, val *uintptr) ZapxLogger
	Uintptrs(key string, val []uintptr) ZapxLogger
	Reflect(key string, val interface{}) ZapxLogger
	Namespace(key string) ZapxLogger
	Stringer(key string, val fmt.Stringer) ZapxLogger
	String(key string, val string) ZapxLogger
	Stringp(key string, val *string) ZapxLogger
	Strings(key string, val []string) ZapxLogger
	Time(key string, val time.Time) ZapxLogger
	Timep(key string, val *time.Time) ZapxLogger
	Times(key string, val []time.Time) ZapxLogger
	Duration(key string, val time.Duration) ZapxLogger
	Durationp(key string, val *time.Duration) ZapxLogger
	Durations(key string, val []time.Duration) ZapxLogger
	Error(error) ZapxLogger
	Object(key string, val zapcore.ObjectMarshaler) ZapxLogger
	Any(key string, value interface{}) ZapxLogger
	Msg(string)
}

type zapx struct {
	fields []zap.Field
	level  zapcore.Level
}

func newZapx(level zapcore.Level) *zapx {
	return &zapx{
		fields: nil,
		level:  level,
	}
}

func (z *zapx) Skip() ZapxLogger {
	z.fields = append(z.fields, zap.Skip())
	return z
}

func (z *zapx) Stack(key string) ZapxLogger {
	z.fields = append(z.fields, zap.Stack(key))
	return z
}

func (z *zapx) StackSkip(key string, skip int) ZapxLogger {
	z.fields = append(z.fields, zap.StackSkip(key, skip))
	return z
}

func (z *zapx) Binary(key string, val []byte) ZapxLogger {
	z.fields = append(z.fields, zap.Binary(key, val))
	return z
}

func (z *zapx) Bool(key string, val bool) ZapxLogger {
	z.fields = append(z.fields, zap.Bool(key, val))
	return z
}

func (z *zapx) Boolp(key string, val *bool) ZapxLogger {
	z.fields = append(z.fields, zap.Boolp(key, val))
	return z
}

func (z *zapx) Bools(key string, val []bool) ZapxLogger {
	z.fields = append(z.fields, zap.Bools(key, val))
	return z
}

func (z *zapx) ByteString(key string, val []byte) ZapxLogger {
	z.fields = append(z.fields, zap.ByteString(key, val))
	return z
}

func (z *zapx) ByteStrings(key string, val [][]byte) ZapxLogger {
	z.fields = append(z.fields, zap.ByteStrings(key, val))
	return z
}

func (z *zapx) Complex128(key string, val complex128) ZapxLogger {
	z.fields = append(z.fields, zap.Complex128(key, val))
	return z
}

func (z *zapx) Complex128p(key string, val *complex128) ZapxLogger {
	z.fields = append(z.fields, zap.Complex128p(key, val))
	return z
}

func (z *zapx) Complex64(key string, val complex64) ZapxLogger {
	z.fields = append(z.fields, zap.Complex64(key, val))
	return z
}

func (z *zapx) Complex64p(key string, val *complex64) ZapxLogger {
	z.fields = append(z.fields, zap.Complex64p(key, val))
	return z
}

func (z *zapx) Float64(key string, val float64) ZapxLogger {
	z.fields = append(z.fields, zap.Float64(key, val))
	return z
}

func (z *zapx) Float64p(key string, val *float64) ZapxLogger {
	z.fields = append(z.fields, zap.Float64p(key, val))
	return z
}

func (z *zapx) Float32(key string, val float32) ZapxLogger {
	z.fields = append(z.fields, zap.Float32(key, val))
	return z
}

func (z *zapx) Float32p(key string, val *float32) ZapxLogger {
	z.fields = append(z.fields, zap.Float32p(key, val))
	return z
}

func (z *zapx) Int(key string, val int) ZapxLogger {
	z.fields = append(z.fields, zap.Int(key, val))
	return z
}

func (z *zapx) Intp(key string, val *int) ZapxLogger {
	z.fields = append(z.fields, zap.Intp(key, val))
	return z
}

func (z *zapx) Int64(key string, val int64) ZapxLogger {
	z.fields = append(z.fields, zap.Int64(key, val))
	return z
}

func (z *zapx) Int64p(key string, val *int64) ZapxLogger {
	z.fields = append(z.fields, zap.Int64p(key, val))
	return z
}

func (z *zapx) Int32(key string, val int32) ZapxLogger {
	z.fields = append(z.fields, zap.Int32(key, val))
	return z
}

func (z *zapx) Int32p(key string, val *int32) ZapxLogger {
	z.fields = append(z.fields, zap.Int32p(key, val))
	return z
}

func (z *zapx) Int16(key string, val int16) ZapxLogger {
	z.fields = append(z.fields, zap.Int16(key, val))
	return z
}

func (z *zapx) Int16p(key string, val *int16) ZapxLogger {
	z.fields = append(z.fields, zap.Int16p(key, val))
	return z
}

func (z *zapx) Int8(key string, val int8) ZapxLogger {
	z.fields = append(z.fields, zap.Int8(key, val))
	return z
}

func (z *zapx) Int8p(key string, val *int8) ZapxLogger {
	z.fields = append(z.fields, zap.Int8p(key, val))
	return z
}

func (z *zapx) Unit(key string, val uint) ZapxLogger {
	z.fields = append(z.fields, zap.Uint(key, val))
	return z
}

func (z *zapx) Unitp(key string, val *uint) ZapxLogger {
	z.fields = append(z.fields, zap.Uintp(key, val))
	return z
}

func (z *zapx) Uint64(key string, val uint64) ZapxLogger {
	z.fields = append(z.fields, zap.Uint64(key, val))
	return z
}

func (z *zapx) Uint64p(key string, val *uint64) ZapxLogger {
	z.fields = append(z.fields, zap.Uint64p(key, val))
	return z
}

func (z *zapx) Uint32(key string, val uint32) ZapxLogger {
	z.fields = append(z.fields, zap.Uint32(key, val))
	return z
}

func (z *zapx) Uint32p(key string, val *uint32) ZapxLogger {
	z.fields = append(z.fields, zap.Uint32p(key, val))
	return z
}

func (z *zapx) Uint16(key string, val uint16) ZapxLogger {
	z.fields = append(z.fields, zap.Uint16(key, val))
	return z
}

func (z *zapx) Uint16p(key string, val *uint16) ZapxLogger {
	z.fields = append(z.fields, zap.Uint16p(key, val))
	return z
}

func (z *zapx) Uint8(key string, val uint8) ZapxLogger {
	z.fields = append(z.fields, zap.Uint8(key, val))
	return z
}

func (z *zapx) Uint8p(key string, val *uint8) ZapxLogger {
	z.fields = append(z.fields, zap.Uint8p(key, val))
	return z
}

func (z *zapx) Uintptr(key string, val uintptr) ZapxLogger {
	z.fields = append(z.fields, zap.Uintptr(key, val))
	return z
}

func (z *zapx) Uintptrp(key string, val *uintptr) ZapxLogger {
	z.fields = append(z.fields, zap.Uintptrp(key, val))
	return z
}

func (z *zapx) Reflect(key string, val interface{}) ZapxLogger {
	z.fields = append(z.fields, zap.Reflect(key, val))
	return z
}

func (z *zapx) Namespace(key string) ZapxLogger {
	z.fields = append(z.fields, zap.Namespace(key))
	return z
}

func (z *zapx) Stringer(key string, val fmt.Stringer) ZapxLogger {
	z.fields = append(z.fields, zap.Stringer(key, val))
	return z
}

func (z *zapx) String(key string, val string) ZapxLogger {
	z.fields = append(z.fields, zap.String(key, val))
	return z
}

func (z *zapx) Stringp(key string, val *string) ZapxLogger {
	z.fields = append(z.fields, zap.Stringp(key, val))
	return z
}

func (z *zapx) Time(key string, val time.Time) ZapxLogger {
	z.fields = append(z.fields, zap.Time(key, val))
	return z
}

func (z *zapx) Timep(key string, val *time.Time) ZapxLogger {
	z.fields = append(z.fields, zap.Timep(key, val))
	return z
}

func (z *zapx) Duration(key string, val time.Duration) ZapxLogger {
	z.fields = append(z.fields, zap.Duration(key, val))
	return z
}

func (z *zapx) Durationp(key string, val *time.Duration) ZapxLogger {
	z.fields = append(z.fields, zap.Durationp(key, val))
	return z
}

func (z *zapx) Array(key string, val zapcore.ArrayMarshaler) ZapxLogger {
	z.fields = append(z.fields, zap.Array(key, val))
	return z
}

func (z *zapx) Complex128s(key string, val []complex128) ZapxLogger {
	z.fields = append(z.fields, zap.Complex128s(key, val))
	return z
}

func (z *zapx) Complex64s(key string, val []complex64) ZapxLogger {
	z.fields = append(z.fields, zap.Complex64s(key, val))
	return z
}

func (z *zapx) Float64s(key string, val []float64) ZapxLogger {
	z.fields = append(z.fields, zap.Float64s(key, val))
	return z
}

func (z *zapx) Float32s(key string, val []float32) ZapxLogger {
	z.fields = append(z.fields, zap.Float32s(key, val))
	return z
}

func (z *zapx) Ints(key string, val []int) ZapxLogger {
	z.fields = append(z.fields, zap.Ints(key, val))
	return z
}

func (z *zapx) Int64s(key string, val []int64) ZapxLogger {
	z.fields = append(z.fields, zap.Int64s(key, val))
	return z
}

func (z *zapx) Int32s(key string, val []int32) ZapxLogger {
	z.fields = append(z.fields, zap.Int32s(key, val))
	return z
}

func (z *zapx) Int16s(key string, val []int16) ZapxLogger {
	z.fields = append(z.fields, zap.Int16s(key, val))
	return z
}

func (z *zapx) Int8s(key string, val []int8) ZapxLogger {
	z.fields = append(z.fields, zap.Int8s(key, val))
	return z
}

func (z *zapx) Units(key string, val []uint) ZapxLogger {
	z.fields = append(z.fields, zap.Uints(key, val))
	return z
}

func (z *zapx) Uint64s(key string, val []uint64) ZapxLogger {
	z.fields = append(z.fields, zap.Uint64s(key, val))
	return z
}

func (z *zapx) Uint32s(key string, val []uint32) ZapxLogger {
	z.fields = append(z.fields, zap.Uint32s(key, val))
	return z
}

func (z *zapx) Uint16s(key string, val []uint16) ZapxLogger {
	z.fields = append(z.fields, zap.Uint16s(key, val))
	return z
}

func (z *zapx) Uint8s(key string, val []uint8) ZapxLogger {
	z.fields = append(z.fields, zap.Uint8s(key, val))
	return z
}

func (z *zapx) Uintptrs(key string, val []uintptr) ZapxLogger {
	z.fields = append(z.fields, zap.Uintptrs(key, val))
	return z
}

func (z *zapx) Strings(key string, val []string) ZapxLogger {
	z.fields = append(z.fields, zap.Strings(key, val))
	return z
}

func (z *zapx) Times(key string, val []time.Time) ZapxLogger {
	z.fields = append(z.fields, zap.Times(key, val))
	return z
}

func (z *zapx) Durations(key string, val []time.Duration) ZapxLogger {
	z.fields = append(z.fields, zap.Durations(key, val))
	return z
}

func (z *zapx) Error(err error) ZapxLogger {
	z.fields = append(z.fields, zap.Error(err))
	return z
}

func (z *zapx) Object(key string, val zapcore.ObjectMarshaler) ZapxLogger {
	z.fields = append(z.fields, zap.Object(key, val))
	return z
}

func (z *zapx) Any(key string, value interface{}) ZapxLogger {
	z.fields = append(z.fields, zap.Any(key, value))
	return z
}

func (z *zapx) Msg(s string) {
	l := _zapx.log.With(z.fields...)
	switch z.level {
	case zapcore.DebugLevel:
		l.Debug(s)
	case zapcore.InfoLevel:
		l.Info(s)
	case zapcore.WarnLevel:
		l.Warn(s)
	case zapcore.ErrorLevel:
		l.Error(s)
	case zapcore.DPanicLevel:
		l.DPanic(s)
	case zapcore.PanicLevel:
		l.Panic(s)
	case zapcore.FatalLevel:
		l.Fatal(s)
	}
}
