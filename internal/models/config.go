package models

type Target struct {
	URLPattern string `json:"url_pattern"`
	Host       string `json:"host"`
	Group      string `json:"group"`
	Timeout    int64  `json:"timeout"`
	Method     string `json:"method"`
}

type Endpoint struct {
	Endpoint string   `json:"endpoint"`
	Method   string   `json:"method"`
	Timeout  string   `json:"timeout"`
	Targets  []Target `json:"targets"`
}

type Config struct {
	Version   int        `json:"version"`
	Endpoints []Endpoint `json:"endpoints"`
}
