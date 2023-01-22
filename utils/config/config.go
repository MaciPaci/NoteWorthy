package config

import (
	"NoteWorthy/assets/env"
	"encoding/json"
	"os"
)

type Config struct {
	Token string `json:"bot_token"`
}

func LoadConfig() (*Config, error) {
	data, err := os.ReadFile(env.ConfigFilePath)
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
