package middleware

import (
	"errors"
	"github.com/Tooooommy/go-one/server/httpx"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/sony/gobreaker"
	"log"
	"net/http"
	"time"
)

type BreakerConfig struct {
	BreakerName     string `json:"breaker_name"`      // 熔断名字
	Timeout         int    `json:"timeout"`           // 请求超时时间
	MaxRequests     int    `json:"max_requests"`      // 最大并发请求
	Interval        int    `json:"interval"`          // 统计周期
	ErrPerThreshold int    `json:"err_per_threshold"` // 允许出现错误比例
	ReqVolThreshold int    `json:"req_vol_threshold"` // 波动期内最小请求数
}

func (b BreakerConfig) GetHystrixCommandConfig() hystrix.CommandConfig {
	return hystrix.CommandConfig{
		Timeout:                b.Timeout,
		MaxConcurrentRequests:  b.MaxRequests,
		SleepWindow:            b.Interval,
		ErrorPercentThreshold:  b.ErrPerThreshold,
		RequestVolumeThreshold: b.ReqVolThreshold,
	}
}

func (b BreakerConfig) GetSonyGoBreaker() gobreaker.Settings {
	return gobreaker.Settings{
		Name:        b.BreakerName,
		MaxRequests: uint32(b.MaxRequests),
		Interval:    time.Duration(b.Interval) * time.Millisecond,
		Timeout:     time.Duration(b.Timeout) * time.Millisecond,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			total := counts.TotalFailures + counts.TotalSuccesses
			return counts.TotalFailures/total*100 > uint32(b.ErrPerThreshold)
		},
		OnStateChange: func(name string, from gobreaker.State, to gobreaker.State) {
			// TODO:
			log.Println(time.Now(), name, from, to)
		},
	}
}

var ErrRequestBreaker = errors.New("request breaker")

func SonyTowStepBreaker(config BreakerConfig) func(http.Handler) http.Handler {
	breaker := gobreaker.NewTwoStepCircuitBreaker(config.GetSonyGoBreaker())
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			promise, err := breaker.Allow()
			if err != nil {
				http.Error(w, ErrRequestBreaker.Error(), http.StatusServiceUnavailable)
			}
			wx := &httpx.CodeResponseWriter{Writer: w}
			defer func() {
				if wx.Code >= http.StatusInternalServerError {
					promise(false)
				}
			}()
			next.ServeHTTP(wx, r)
		})
	}
}

func SonyBreaker(config BreakerConfig) func(http.Handler) http.Handler {
	breaker := gobreaker.NewCircuitBreaker(config.GetSonyGoBreaker())
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := breaker.Execute(func() (interface{}, error) {
				var err error
				wx := &httpx.CodeResponseWriter{Writer: w}
				defer func() {
					if wx.Code >= http.StatusInternalServerError {
						err = ErrRequestBreaker
					}
				}()
				next.ServeHTTP(wx, r)
				return nil, err
			})
			if err != nil {
				http.Error(w, ErrRequestBreaker.Error(), http.StatusServiceUnavailable)
			}
		})
	}
}

func HystrixBreaker(config BreakerConfig) func(http.Handler) http.Handler {
	hystrix.ConfigureCommand(config.BreakerName, config.GetHystrixCommandConfig())
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			err := hystrix.Do(config.BreakerName, func() error {
				wx := &httpx.CodeResponseWriter{Writer: w}
				var err error
				defer func() {
					if wx.Code >= http.StatusInternalServerError {
						err = ErrRequestBreaker
					}
				}()
				next.ServeHTTP(wx, r)
				return err
			}, nil)
			if err != nil {
				http.Error(w, ErrRequestBreaker.Error(), http.StatusServiceUnavailable)
			}
		})
	}
}
