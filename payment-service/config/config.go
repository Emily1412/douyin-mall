package config

import (
	"log"

	"douyin-mall/payment-service/internal/model"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// LoadConfig 加载配置文件
func LoadConfig() error {
	viper.SetConfigFile("config/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

// InitDB 初始化数据库连接和自动迁移
func InitDB() {
	var err error
	dsn := viper.GetString("mysql.user") + ":" +
		viper.GetString("mysql.password") + "@tcp(" +
		viper.GetString("mysql.host") + ":" +
		viper.GetString("mysql.port") + ")/" +
		viper.GetString("mysql.database") + "?charset=utf8&parseTime=True&loc=Local"

	// 连接数据库
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// 自动迁移 payment 表
	if err := DB.AutoMigrate(&model.Payment{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	log.Println("Database migrated")
}
