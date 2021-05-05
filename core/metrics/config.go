package metrics

type Config struct {
	Namespace   string              `json:"namespace"`
	Name        string              `json:"name"`
	Subsystem   string              `json:"subsystem"`
	Help        string              `json:"help"`
	ConstLabels map[string]string   `json:"const_labels"`
	Buckets     []float64           `json:"buckets"`
	Objectives  map[float64]float64 `json:"objectives"`
	MaxAge      int64               `json:"max_age"`
	AgeBuckets  uint32              `json:"age_buckets"`
	BufCap      uint32              `json:"buf_cap"`
}
