package discov

type Config struct {
	Name     string   `json:"name"`      // key
	Hosts    []string `json:"endpoints"` // val
	Username string   `json:"username"`
	Password string   `json:"password"`
}

func (c Config) HaveEtcd() bool {
	return len(c.Hosts) > 0 && len(c.Name) > 0
}
