package breaker

import (
	"github.com/Tooooommy/go-one/core/zapx"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/sony/gobreaker"
	"time"
)

type Config struct {
	Name            string `json:"name"`              // 熔断名字
	Timeout         int    `json:"timeout"`           // 请求超时时间
	MaxRequests     int    `json:"max_requests"`      // 最大并发请求
	Interval        int    `json:"interval"`          // 统计周期，单位微秒
	ErrPerThreshold int    `json:"err_per_threshold"` // 允许出现错误比例
	ReqVolThreshold int    `json:"req_vol_threshold"` // 波动期内最小请求数
}

func readyToTrip(errPerThreshold int) func(counts gobreaker.Counts) bool {
	if errPerThreshold > 0 {
		return func(counts gobreaker.Counts) bool {
			total := counts.TotalFailures + counts.TotalSuccesses
			return counts.TotalFailures/total*100 > uint32(errPerThreshold)
		}
	}
	return nil
}

func (cfg Config) GetGoBreakerSettings() gobreaker.Settings {
	return gobreaker.Settings{
		Name:        cfg.Name,
		MaxRequests: uint32(cfg.MaxRequests),
		Interval:    time.Duration(cfg.Interval) * time.Millisecond,
		Timeout:     time.Duration(cfg.Timeout) * time.Millisecond,
		ReadyToTrip: readyToTrip(cfg.ErrPerThreshold),
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			zapx.Info().String("name", name).
				Int("from_state", int(from)).
				Int("to_state", int(to)).
				Msg("状态发生变化")
		},
	}
}

func (cfg Config) GetHystrixConfig() hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                cfg.Timeout,
		MaxConcurrentRequests:  cfg.MaxRequests,
		SleepWindow:            cfg.Interval,
		ErrorPercentThreshold:  cfg.ErrPerThreshold,
		RequestVolumeThreshold: cfg.ReqVolThreshold,
	}
}
