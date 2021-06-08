package breaker

import (
	"context"
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/sony/gobreaker"
	"time"
)

type Conf struct {
	Name        string        `json:"name"`         // 熔断名字
	MaxRequests int           `json:"max_requests"` // 最大并发请求
	Interval    time.Duration `json:"interval"`     // 统计周期
	Timeout     time.Duration `json:"timeout"`      // 请求超时时间
	ErrPercent  int           `json:"err_percent"`  // 允许出现错误比例
}

// TODO: 默认
func DefaultConf() *Conf {
	return &Conf{
		Name:        "go-one",
		Timeout:     60,
		MaxRequests: 1024,
		Interval:    30,
		ErrPercent:  60,
	}
}

func readyToTrip(errPercent int) func(counts gobreaker.Counts) bool {
	if errPercent > 0 {
		return func(counts gobreaker.Counts) bool {
			return counts.TotalFailures/counts.Requests*100 > uint32(errPercent)
		}
	}
	return nil
}

func onStateChange(name string, from gobreaker.State, to gobreaker.State) {
	zapx.Info(context.Background()).String("breaker_name", name).
		Int("from_state", int(from)).Int("to_state", int(to)).
		Msg("状态发生变化")
}

func (cfg *Conf) GetGoBreakerSettings() gobreaker.Settings {
	return gobreaker.Settings{
		Name:          cfg.Name,
		MaxRequests:   uint32(cfg.MaxRequests),
		Interval:      cfg.Interval,
		Timeout:       cfg.Timeout,
		ReadyToTrip:   readyToTrip(cfg.ErrPercent),
		OnStateChange: onStateChange,
	}
}
