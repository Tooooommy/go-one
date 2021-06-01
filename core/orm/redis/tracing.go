package redis

import (
	"context"
	"github.com/go-redis/redis/extra/rediscmd/v8"
	"github.com/go-redis/redis/v8"
	"github.com/opentracing/opentracing-go"
)

type TracingHook struct{}

func (*TracingHook) BeforeProcess(ctx context.Context, cmd redis.Cmder) (context.Context, error) {
	span := opentracing.SpanFromContext(ctx)
	name := "redis." + cmd.FullName()
	if span == nil {
		span = opentracing.StartSpan(name)
	} else {
		span, ctx = opentracing.StartSpanFromContext(ctx, name)
	}
	span.SetTag("db.system", "redis")
	span.SetTag("db.statement", rediscmd.CmdString(cmd))
	return ctx, nil
}

func (*TracingHook) AfterProcess(ctx context.Context, cmd redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("error", cmd.Err())
	span.Finish()
	return nil
}

func (*TracingHook) BeforeProcessPipeline(ctx context.Context, cmds []redis.Cmder) (context.Context, error) {
	span := opentracing.SpanFromContext(ctx)
	summary, cmdsString := rediscmd.CmdsString(cmds)
	name := "redis.pipeline." + summary
	if span == nil {
		span = opentracing.StartSpan(name)
	} else {
		span, ctx = opentracing.StartSpanFromContext(ctx, name)
	}
	span.SetTag("db.system", "redis")
	span.SetTag("db.redis.num_cmd", len(cmds))
	span.SetTag("db.statement", cmdsString)
	return ctx, nil
}

func (TracingHook) AfterProcessPipeline(ctx context.Context, cmds []redis.Cmder) error {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("error", cmds[0].Err())
	span.Finish()
	return nil
}
