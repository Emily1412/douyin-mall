package model

import "gorm.io/gorm"

// Order 订单模型
type Order struct {
	gorm.Model
	UserID       uint32 `gorm:"index"`
	UserCurrency string
	Address      string // 简单存储为字符串（可以使用 JSON 格式存储复杂结构）
	Email        string
	Status       string      // 例如 "pending", "paid" 等状态
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID"`
}

// OrderItem 订单项模型
type OrderItem struct {
	gorm.Model
	OrderID   uint
	ProductID uint32
	Quantity  int32
	Cost      float64
}
