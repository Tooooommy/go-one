package config

type BreakerConfig struct {
	Name            string `json:"name"`              // 熔断名字
	Timeout         int    `json:"timeout"`           // 请求超时时间
	MaxRequests     int    `json:"max_requests"`      // 最大并发请求
	Interval        int    `json:"interval"`          // 统计周期，单位微秒
	ErrPerThreshold int    `json:"err_per_threshold"` // 允许出现错误比例
	ReqVolThreshold int    `json:"req_vol_threshold"` // 波动期内最小请求数
}
