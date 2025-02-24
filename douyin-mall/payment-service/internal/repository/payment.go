
// 支付仓库
package repository

import (
	"context"
	"douyin-mall/payment-service/model"
	"douyin-mall/pkg/db"
)

type PaymentRepository struct{}

// 创建支付仓库
func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{}
}

// 创建交易
func (r *PaymentRepository) CreateTransaction(ctx context.Context, tx *model.Transaction) error {
	return db.DB.WithContext(ctx).Create(tx).Error
}

// 更新交易状态
func (r *PaymentRepository) UpdateTransactionStatus(ctx context.Context, transactionID, status string) error {
	return db.DB.WithContext(ctx).Model(&model.Transaction{}).
		Where("transaction_id = ?", transactionID).
		Update("status", status).Error
}

func (r *PaymentRepository) GetTransactionByID(ctx context.Context, transactionID string) (*model.Transaction, error) {
	var tx model.Transaction
	err := db.DB.WithContext(ctx).Where("transaction_id = ?", transactionID).First(&tx).Error
	return &tx, err
}

func (r *PaymentRepository) CreatePaymentLog(ctx context.Context, log *model.PaymentLog) error {
	return db.DB.WithContext(ctx).Create(log).Error
}
