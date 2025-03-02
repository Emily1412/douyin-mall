package repository

import (
	"douyin-mall/payment-service/internal/model"

	"gorm.io/gorm"
)

// PaymentRepository 定义了与支付相关的数据库操作
type PaymentRepository struct {
	DB *gorm.DB
}

// NewPaymentRepository 创建一个新的 PaymentRepository
func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{DB: db}
}

// CreatePaymentRecord 创建支付记录
func (r *PaymentRepository) CreatePaymentRecord(payment *model.Payment) error {
	return r.DB.Create(payment).Error
}

// GetPaymentRecordByTransactionID 根据交易 ID 获取支付记录
func (r *PaymentRepository) GetPaymentRecordByTransactionID(transactionID string) (*model.Payment, error) {
	var payment model.Payment
	if err := r.DB.Where("transaction_id = ?", transactionID).First(&payment).Error; err != nil {
		return nil, err
	}
	return &payment, nil
}
