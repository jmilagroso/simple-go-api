package models

// Index struct
type Index struct {
	ServerTime   string `redis:"server_time" json:"server_time"`
	GoVersion    string `redis:"go_version" json:"go_version"`
	CacheTimeout int16  `redis:"cache_timeout" json:"cache_timeout"`
}
