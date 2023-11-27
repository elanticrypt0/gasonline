package access

import "github.com/elanticrypt0/go4it"

type AccessConfig struct {
	IsEnabled bool `json:"is_enabled" toml:"is_enabled"`
}

func LoadConfig(config *AccessConfig) {
	go4it.ReadAndParseToml("./config/access.toml", &config)
}
