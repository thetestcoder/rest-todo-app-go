package configs

import (
	"encoding/json"
	"os"
)

type Config struct {
	Storage struct {
		FilePath string `json:"file_path"`
	}
}

func NewConfig(path string) *Config {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil
	}

	if config.Storage.FilePath == "" {
		config.Storage.FilePath = "tasks.json"
	}

	return &config
}
