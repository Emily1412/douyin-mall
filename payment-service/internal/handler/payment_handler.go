package handler

import (
	"context"
	"douyin-mall/payment-service/internal/service"
	"douyin-mall/payment-service/proto"
)

// PaymentHandler 处理支付请求的 gRPC 处理器
type PaymentHandler struct {
	proto.UnimplementedPaymentServiceServer // 确保嵌入 UnimplementedPaymentServiceServer
	PaymentService                          *service.PaymentService
}

// NewPaymentHandler 创建新的 PaymentHandler
func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{PaymentService: paymentService}
}

// Pay 处理支付请求
func (h *PaymentHandler) Pay(ctx context.Context, req *proto.PayReq) (*proto.PayResp, error) {
	// 调用业务层的 Pay 方法处理支付
	resp, err := h.PaymentService.Pay(req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
