package config

import (
	"github.com/pelletier/go-toml"
)

type DatabaseConfig struct {
	Host   string `toml:"host"`
	Port   string `toml:"port"`
	User   string `toml:"user"`
	DbName string `toml:"dbName"`
}

type ServerConfig struct {
	BindAddr string `toml:"bind_addr"`
}

type Config struct {
	Database DatabaseConfig `toml:"database"`
	Server   ServerConfig   `toml:"server"`
}

func NewConfig() *Config {
	return &Config{
		DatabaseConfig{
			"localhost",
			"5432",
			"postgres",
			"graph_db",
		},
		ServerConfig{
			":8080",
		},
	}
}

func LoadConfig(filePath string) (*Config, error) {
	config := NewConfig()
	file, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	if err = file.Unmarshal(config); err != nil {
		return nil, err
	}

	return config, nil
}
