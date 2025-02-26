package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type Config struct {
	Database struct {
		DSN string `mapstructure:"dsn"`
	} `mapstructure:"database"`
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
}

func LoadConfig() (*Config, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("failed to read config file: %w", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config: %w", err)
	}

	return &cfg, nil
}
