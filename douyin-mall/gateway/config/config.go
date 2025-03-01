package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Mode string `yaml:"mode"`
	} `yaml:"server"`

	Services struct {
		Payment  string `yaml:"payment"`
		User     string `yaml:"user"`
		Order    string `yaml:"order"`
		Product  string `yaml:"product"`
		Cart     string `yaml:"cart"`
		Checkout string `yaml:"checkout"`
	} `yaml:"services"`

	JWT struct {
		Secret     string `yaml:"secret"`
		ExpireTime int    `yaml:"expireTime"` // 单位：小时
	} `yaml:"jwt"`

	Redis struct {
		Addr     string `yaml:"addr"`
		Password string `yaml:"password"`
		DB       int    `yaml:"db"`
	} `yaml:"redis"`

	PaymentService struct {
		Endpoint string `yaml:"endpoint"`
	} `yaml:"payment_service"`
}

var GlobalConfig Config

func Init() error {
	// 从环境变量获取配置文件路径
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "D:\\headlight\\douyin-mall\\douyin-mall\\gateway\\config\\gateway.yaml"
	}

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	// 解析配置
	return yaml.Unmarshal(data, &GlobalConfig)
}
