package pulgin

import (
	"github.com/opentracing/opentracing-go"
	"gorm.io/gorm"
	"gorm.io/plugin/opentracing"
)

func NewTracer() gorm.Plugin {
	return gormopentracing.New(
		gormopentracing.WithTracer(opentracing.GlobalTracer()),
		gormopentracing.WithLogResult(true),
		gormopentracing.WithSqlParameters(true),
	)
}
