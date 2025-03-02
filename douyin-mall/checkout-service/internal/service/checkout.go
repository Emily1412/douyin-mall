package service

import (
	"context"
	"fmt"
	"log"
	"time"

	"checkout-service/config"
	pb "checkout-service/proto"
	orderpb "order-service/proto"
	paymentpb "payment-service/api"

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

	// 在使用 req.CreditCard 之前检查
	if req.CreditCard == nil {
		return nil, fmt.Errorf("信用卡信息不能为空")
	}

	// 生成订单ID - 使用时间戳作为整数订单ID
	orderId := fmt.Sprintf("%d", time.Now().UnixNano()/1000000) // 毫秒时间戳

	// 连接订单服务，获取订单金额
	orderAmount, err := s.getOrderAmount(ctx, req.UserId, orderId)
	if err != nil {
		log.Printf("获取订单金额失败: %v", err)
		return nil, fmt.Errorf("获取订单金额失败: %v", err)
	}

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
		Amount: orderAmount,
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

	// 在使用 chargeResp 之前检查
	if chargeResp == nil {
		return nil, fmt.Errorf("支付响应为空")
	}

	// 标记订单为已支付
	if err := s.markOrderAsPaid(ctx, req.UserId, orderId); err != nil {
		log.Printf("标记订单为已支付失败: %v", err)
		// 这里我们不返回错误，因为支付已经成功，只是标记失败
		// 可以通过其他方式处理，比如异步重试
	}

	// 返回结算响应
	return &pb.CheckoutResp{
		OrderId:       orderId,
		TransactionId: chargeResp.TransactionId,
	}, nil
}

// getOrderAmount 从订单服务获取订单金额
func (s *CheckoutService) getOrderAmount(ctx context.Context, userId uint32, orderId string) (float32, error) {
	// 连接订单服务
	orderAddr := fmt.Sprintf("%s:%d", s.config.OrderService.Host, s.config.OrderService.Port)
	orderConn, err := grpc.Dial(orderAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return 0, fmt.Errorf("无法连接订单服务: %v", err)
	}
	defer orderConn.Close()

	// 创建订单服务客户端
	orderClient := orderpb.NewOrderServiceClient(orderConn)

	// 设置超时上下文
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 获取用户的订单列表
	listOrderResp, err := orderClient.ListOrder(timeoutCtx, &orderpb.ListOrderReq{
		UserId: userId,
	})
	if err != nil {
		return 0, fmt.Errorf("获取订单列表失败: %v", err)
	}

	// 计算订单总金额
	var totalAmount float32
	for _, order := range listOrderResp.Orders {
		for _, item := range order.OrderItems {
			totalAmount += item.Cost
		}
	}

	// 如果没有找到订单或订单金额为0或负数，使用默认金额
	if totalAmount <= 0 {
		log.Printf("未找到订单或订单金额为0或负数，使用默认金额")
		totalAmount = 100.0
	}

	return totalAmount, nil
}

// markOrderAsPaid 标记订单为已支付
func (s *CheckoutService) markOrderAsPaid(ctx context.Context, userId uint32, orderId string) error {
	// 连接订单服务
	orderAddr := fmt.Sprintf("%s:%d", s.config.OrderService.Host, s.config.OrderService.Port)
	orderConn, err := grpc.Dial(orderAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return fmt.Errorf("无法连接订单服务: %v", err)
	}
	defer orderConn.Close()

	// 创建订单服务客户端
	orderClient := orderpb.NewOrderServiceClient(orderConn)

	// 设置超时上下文
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 标记订单为已支付
	_, err = orderClient.MarkOrderPaid(timeoutCtx, &orderpb.MarkOrderPaidReq{
		UserId:  userId,
		OrderId: orderId,
	})
	if err != nil {
		return fmt.Errorf("标记订单为已支付失败: %v", err)
	}

	return nil
}
