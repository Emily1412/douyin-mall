// 支付仓库
package repository

import (
	"context"
	"douyin-mall/payment-service/internal/model"
	"douyin-mall/payment-service/pkg/database"

	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

// 创建支付仓库
func NewPaymentRepository() *PaymentRepository {
	return &PaymentRepository{
		db: database.DB,
	}
}

// 创建交易
func (r *PaymentRepository) CreateTransaction(ctx context.Context, tx *model.Transaction) error {
	return r.db.WithContext(ctx).Create(tx).Error
}

// 更新交易状态
func (r *PaymentRepository) UpdateTransactionStatus(ctx context.Context, transactionID, status string) error {
	return r.db.WithContext(ctx).Model(&model.Transaction{}).
		Where("transaction_id = ?", transactionID).
		Update("status", status).Error
}

func (r *PaymentRepository) GetTransactionByID(ctx context.Context, transactionID string) (*model.Transaction, error) {
	var tx model.Transaction
	err := r.db.WithContext(ctx).Where("transaction_id = ?", transactionID).First(&tx).Error
	return &tx, err
}

func (r *PaymentRepository) CreatePaymentLog(ctx context.Context, log *model.PaymentLog) error {
	return r.db.WithContext(ctx).Create(log).Error
}
