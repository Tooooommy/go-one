package config

type JwtConfig struct {
	Secret      string `json:"secret"`
	Timeout     int    `json:"timeout"` // hour
	PreSecret   string `json:"pre_secret"`
	CanRefresh  bool   `json:"can_refresh"` // 可以刷新不
	ValidRealIp bool   `json:"valid_real_ip"`
	ValidDevice bool   `json:"valid_device"`
}
