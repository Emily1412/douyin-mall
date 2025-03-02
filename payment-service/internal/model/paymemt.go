package model

// Payment 定义支付记录模型
type Payment struct {
	ID            uint   `gorm:"primaryKey"`
	TransactionID string `gorm:"unique"`
	OrderID       uint   `gorm:"index"`
	Amount        float64
	Status        string
	CreatedAt     int64
	UpdatedAt     int64
}
