package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MysqlConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

type ConfigType struct {
	Mysql MysqlConfig `yaml:"mysql"`
}

var Config ConfigType
var DB *gorm.DB

// LoadConfig 从指定路径加载 YAML 配置文件
func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(data, &Config)
}

// InitDB 初始化 MySQL 数据库连接
func InitDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		Config.Mysql.User,
		Config.Mysql.Password,
		Config.Mysql.Host,
		Config.Mysql.Port,
		Config.Mysql.Database,
	)
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	log.Println("Connected to MySQL")
}
