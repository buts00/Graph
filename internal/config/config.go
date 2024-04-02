package config

import (
	"github.com/spf13/viper"
	"os"
)

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
}

type ServerConfig struct {
	BindAddr string
}

type Config struct {
	Database DatabaseConfig
	Server   ServerConfig
}

func LoadConfig(configPath string) (*Config, error) {
	viper.AddConfigPath(configPath)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		Database: DatabaseConfig{
			viper.GetString("db.host"),
			viper.GetString("db.port"),
			viper.GetString("db.user"),
			os.Getenv("DB_PASSWORD"),
			viper.GetString("db.db_name"),
		},
		Server: ServerConfig{
			viper.GetString("bind_addr"),
		},
	}, nil
}
