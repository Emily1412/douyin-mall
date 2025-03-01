package model

import "gorm.io/gorm"

// CartItem 购物车项模型
type CartItem struct {
	gorm.Model
	ProductID uint32 `gorm:"index"`
	Quantity  int32  `gorm:"not null"`
	CartID    uint   // 修改为 uint，与 Cart.ID 保持一致
}

// Cart 购物车模型
type Cart struct {
	gorm.Model
	UserID uint32     `gorm:"index;unique"`
	Items  []CartItem `gorm:"foreignKey:CartID"`
}
