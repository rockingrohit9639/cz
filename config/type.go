package config

type Format struct {
	Name    string `json:"name"`
	Pattern string `json:"pattern"`
}

type Config struct {
	Formats []Format `json:"formats"`
	Default string   `json:"default"`
}
