package model

import (
	"time"

	"gorm.io/gorm"
)

// Transaction 支付交易记录
type Transaction struct {
	TransactionID string `gorm:"primaryKey;type:varchar(64)"`
	OrderID       string `gorm:"index;type:varchar(64)"`
	UserID        uint32 `gorm:"index"`
	Amount        float64
	Status        string `gorm:"type:varchar(20)"` // PENDING, SUCCESS, FAILED
	PaymentMethod string `gorm:"type:varchar(20)"` // CREDIT_CARD, ALIPAY, WECHAT etc.
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
}

// PaymentLog 支付日志
type PaymentLog struct {
	gorm.Model
	TransactionID string `gorm:"index;type:varchar(64)"`
	PaymentMethod string `gorm:"type:varchar(20)"`
	PaymentStatus string `gorm:"type:varchar(20)"`
	ErrorMessage  string `gorm:"type:text"`
}