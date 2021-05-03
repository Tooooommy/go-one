package config

type HttpConfig struct {
	Name         string `json:"name"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	CertFile     string `json:"cert_file"`
	KeyFile      string `json:"key_file"`
	MaxConns     int    `json:"max_conns"`
	MaxBytes     int64  `json:"max_bytes"`
	Timeout      int64  `json:"timeout"`
	CpuThreshold int64  `json:"cpu_threshold"`
}
