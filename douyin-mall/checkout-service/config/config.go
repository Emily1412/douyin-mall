package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// Config 结构体定义了应用程序的配置
type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	} `yaml:"server"`

	PaymentService struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	} `yaml:"payment_service"`
}

// LoadConfig 从配置文件加载配置
func LoadConfig() (*Config, error) {
	// 获取当前工作目录
	wd, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("获取工作目录失败: %v", err)
	}

	// 构建配置文件路径
	configPath := filepath.Join(wd, "config", "config.yaml")

	// 读取配置文件
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("读取配置文件失败: %v", err)
	}

	// 解析配置
	var config Config
	if err := yaml.Unmarshal(data, &config); err != nil {
		return nil, fmt.Errorf("解析配置文件失败: %v", err)
	}

	log.Printf("配置加载成功: %+v", config)
	return &config, nil
}
