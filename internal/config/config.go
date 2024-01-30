package config

import (
	"github.com/pelletier/go-toml"
)

type DatabaseConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	DbName   string `toml:"dbName"`
}

type Config struct {
	Database DatabaseConfig `toml:"database"`
}

func NewConfig() *Config {
	return &Config{}
}

func LoadConfig(filePath string) (*Config, error) {
	config := NewConfig()
	file, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err := file.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
