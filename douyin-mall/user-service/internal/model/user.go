package model

import (
	"time"

	"gorm.io/gorm"
)

// Base 基础模型
type Base struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// User 用户模型
type User struct {
	Base
	Email    string `gorm:"unique;not null"` // 用户邮箱，设置为唯一
	Password string `gorm:"not null"`        // 用户密码
}

// TableName 设置表名
func (u User) TableName() string {
	return "user"
}

// GetByEmail 通过邮箱查询用户
func GetByEmail(db *gorm.DB, email string) (user *User, err error) {
	err = db.Where("email = ?", email).First(&user).Error
	return
}

// Create 创建用户
func Create(db *gorm.DB, user *User) error {
	return db.Create(user).Error
}
