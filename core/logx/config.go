package logx

type mode string

const (
	StdMode mode = "std"
	LogMode mode = "log"
	ELKMode mode = "elk"
)

type Config struct {
	Name  string `json:"name"`
	Level int    `json:"level"`
}

type StdModeConfig struct {
	Config
}

type LogModeConfig struct {
	Config
	Filename   string `json:"filename"`
	MaxSize    int    `json:"max_size"`
	MaxAge     int    `json:"max_age"`
	MaxBackups int    `json:"max_backups"`
	LocalTime  bool   `json:"local_time"`
	Compress   bool   `json:"compress"`
}

type ElkModeConfig struct {
	Config
}
