package config

import (
	"encoding/json"
	"os"
)

// Config is a container for configuration variables
type Config struct {
	Token string `json:"bot_token"`
}

// LoadConfig loads configuration from JSON file into Conf struct
func LoadConfig(filepath string) (*Config, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	var conf Config
	err = json.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return &conf, nil
}
