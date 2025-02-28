package database

import (
	"fmt"

	"douyin-mall/payment-service/internal/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DBName   string
}

var DB *gorm.DB

func InitWithConfig(config Config) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Username,
		config.Password,
		config.Host,
		config.Port,
		config.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db
	return nil
}

func TestConnection() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Ping()
}

func AutoMigrate() error {
	return DB.AutoMigrate(
		&model.Transaction{},
		&model.PaymentLog{},
	)
}
