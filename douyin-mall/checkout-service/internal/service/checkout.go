package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"checkout-service/config"
	pb "checkout-service/proto"
	paymentpb "payment-service/api"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// CheckoutService 实现了 CheckoutServiceServer 接口
type CheckoutService struct {
	pb.UnimplementedCheckoutServiceServer
	config *config.Config
}

// NewCheckoutService 创建一个新的结算服务实例
func NewCheckoutService(cfg *config.Config) *CheckoutService {
	return &CheckoutService{
		config: cfg,
	}
}

// Checkout 处理结算请求
func (s *CheckoutService) Checkout(ctx context.Context, req *pb.CheckoutReq) (*pb.CheckoutResp, error) {
	log.Printf("收到结算请求: %v", req)

	// 生成订单ID
	orderId := uuid.New().String()

	// 连接支付服务
	paymentAddr := fmt.Sprintf("%s:%d", s.config.PaymentService.Host, s.config.PaymentService.Port)
	paymentConn, err := grpc.Dial(paymentAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Printf("无法连接支付服务: %v", err)
		return nil, fmt.Errorf("无法连接支付服务: %v", err)
	}
	defer paymentConn.Close()

	// 创建支付服务客户端
	paymentClient := paymentpb.NewPaymentServiceClient(paymentConn)

	// 准备支付请求
	chargeReq := &paymentpb.ChargeReq{
		Amount: 100.0, // 这里应该从订单中获取实际金额
		CreditCard: &paymentpb.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.CreditCardNumber,
			CreditCardCvv:             req.CreditCard.CreditCardCvv,
			CreditCardExpirationYear:  req.CreditCard.CreditCardExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.CreditCardExpirationMonth,
		},
		OrderId: orderId,
		UserId:  req.UserId,
	}

	// 设置超时上下文
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 调用支付服务
	chargeResp, err := paymentClient.Charge(timeoutCtx, chargeReq)
	if err != nil {
		log.Printf("支付处理失败: %v", err)
		return nil, fmt.Errorf("支付处理失败: %v", err)
	}

	// 返回结算响应
	return &pb.CheckoutResp{
		OrderId:       orderId,
		TransactionId: chargeResp.TransactionId,
	}, nil
}
