package config

import (
	"encoding/json"
	"os"
)

var (
	configPath = "config.json"
)

type Config struct {
	Addr string `json:"addr"`
}

func NewConfig() (*Config, error) {
	data, err := os.ReadFile("config.json")
	if err != nil {
		return nil, err
	}
	newConfig := &Config{}
	if err = json.Unmarshal(data, newConfig); err != nil {
		return nil, err
	}
	return newConfig, nil
}
