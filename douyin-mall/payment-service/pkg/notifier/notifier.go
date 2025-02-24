
// 支付结果通知
package payment

import (
	"context"
	"douyin-mall/payment-service/model"
	"fmt"
)

// PaymentNotifier 支付结果通知接口
type PaymentNotifier interface {
	NotifySuccess(ctx context.Context, transaction *model.Transaction) error
	NotifyFailure(ctx context.Context, transaction *model.Transaction, err error) error
}

// DefaultNotifier 默认通知实现
type DefaultNotifier struct{}

func NewDefaultNotifier() *DefaultNotifier {
	return &DefaultNotifier{}
}

func (n *DefaultNotifier) NotifySuccess(ctx context.Context, transaction *model.Transaction) error {
	// TODO: 实现实际的通知逻辑
	// 1. 发送消息到消息队列
	// 2. 调用订单服务更新订单状态
	// 3. 发送邮件或短信通知用户
	fmt.Printf("Payment success notification: TransactionID=%s, Amount=%.2f\n",
		transaction.TransactionID, transaction.Amount)
	return nil
}

func (n *DefaultNotifier) NotifyFailure(ctx context.Context, transaction *model.Transaction, err error) error {
	// TODO: 实现实际的通知逻辑
	fmt.Printf("Payment failure notification: TransactionID=%s, Error=%v\n",
		transaction.TransactionID, err)
	return nil
}
