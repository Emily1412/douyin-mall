package service

import (
	"douyin-mall/payment-service/internal/model"
	"douyin-mall/payment-service/internal/repository"
	"douyin-mall/payment-service/proto"
	"fmt"
)

// PaymentService 定义了支付服务
type PaymentService struct {
	Repo *repository.PaymentRepository
}

// NewPaymentService 创建新的 PaymentService
func NewPaymentService(repo *repository.PaymentRepository) *PaymentService {
	return &PaymentService{Repo: repo}
}

// Pay 处理支付请求
func (s *PaymentService) Pay(req *proto.PayReq) (*proto.PayResp, error) {
	// 模拟支付处理
	transactionID := fmt.Sprintf("txn_%d", req.OrderId)
	payment := &model.Payment{
		TransactionID: transactionID,
		OrderID:       uint(req.OrderId),
		Amount:        float64(req.Amount),
		Status:        "success", // 你可以根据实际逻辑更改状态
	}

	// 保存支付记录
	if err := s.Repo.CreatePaymentRecord(payment); err != nil {
		return nil, err
	}

	// 返回支付响应
	return &proto.PayResp{
		TransactionId: payment.TransactionID,
		Status:        payment.Status,
	}, nil
}
