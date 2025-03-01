package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`

	Database struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		DBName   string `yaml:"dbname"`
	} `yaml:"database"`
}

var GlobalConfig Config

func Init() error {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "D:\\headlight\\douyin-mall\\douyin-mall\\payment-service\\config\\config.yaml"
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(data, &GlobalConfig)
}
