package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type (
	Config struct {
		DB    `yaml:"datasource"`
		Log   `yaml:"log"`
		Files `yaml:"files"`
	}

	DB struct {
		URL string `yaml:"url"`
	}

	Log struct {
		Level string `yaml:"level"`
	}

	Files struct {
		Path string `yaml:"path"`
	}
)

func Load(path string) (*Config, error) {
	var cfg Config
	err := cleanenv.ReadConfig(path, &cfg)
	if err != nil {
		return nil, fmt.Errorf("error reading config file: %s", err)
	}
	err = cleanenv.UpdateEnv(&cfg)
	if err != nil {
		return nil, fmt.Errorf("error updating env: %s", err)
	}
	return &cfg, nil
}
