package service

import (
	"context"
	"payment-service/internal/model"
	"payment-service/internal/repository"
	payment "payment-service/pkg/processor"

	//proto所在地
	pb "payment-service/api"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PaymentServer struct {
	pb.UnimplementedPaymentServiceServer
	repo      *repository.PaymentRepository
	processor *payment.CreditCardProcessor
}

func NewPaymentServer() *PaymentServer {
	return &PaymentServer{
		repo:      repository.NewPaymentRepository(),
		processor: payment.NewCreditCardProcessor(),
	}
}

func (s *PaymentServer) Charge(ctx context.Context, req *pb.ChargeReq) (*pb.ChargeResp, error) {
	log.Printf("Payment request received: %v", req)

	if err := validateChargeRequest(req); err != nil {
		log.Printf("Validation failed: %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	// 2. 创建交易记录
	transactionID := uuid.New().String()
	tx := &model.Transaction{
		TransactionID: transactionID,
		OrderID:       req.OrderId,
		UserID:        req.UserId,
		Amount:        float64(req.Amount),
		Status:        "PENDING",
		PaymentMethod: "CREDIT_CARD",
	}

	if err := s.repo.CreateTransaction(ctx, tx); err != nil {
		return nil, status.Error(codes.Internal, "Failed to create transaction")
	}

	// 3. 处理支付
	if err := s.processPayment(ctx, req.CreditCard, tx); err != nil {
		// 更新交易状态为失败
		s.repo.UpdateTransactionStatus(ctx, transactionID, "FAILED")
		// 记录失败日志
		s.repo.CreatePaymentLog(ctx, &model.PaymentLog{
			TransactionID: transactionID,
			PaymentMethod: "CREDIT_CARD",
			PaymentStatus: "FAILED",
			ErrorMessage:  err.Error(),
		})
		return nil, status.Error(codes.Internal, "Payment processing failed")
	}

	// 4. 更新交易状态为成功
	if err := s.repo.UpdateTransactionStatus(ctx, transactionID, "SUCCESS"); err != nil {
		return nil, status.Error(codes.Internal, "Failed to update transaction status")
	}

	// 5. 记录成功日志
	s.repo.CreatePaymentLog(ctx, &model.PaymentLog{
		TransactionID: transactionID,
		PaymentMethod: "CREDIT_CARD",
		PaymentStatus: "SUCCESS",
	})

	return &pb.ChargeResp{
		TransactionId: transactionID,
	}, nil
}

func validateChargeRequest(req *pb.ChargeReq) error {
	if req.Amount <= 0 {
		return fmt.Errorf("invalid amount")
	}
	if req.OrderId == "" {
		return fmt.Errorf("order ID is required")
	}
	if req.UserId == 0 {
		return fmt.Errorf("user ID is required")
	}
	if req.CreditCard == nil {
		return fmt.Errorf("credit card information is required")
	}
	return nil
}

func (s *PaymentServer) processPayment(ctx context.Context, card *pb.CreditCardInfo, tx *model.Transaction) error {
	// 添加超时控制
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// 使用支付处理器处理支付
	err := s.processor.Process(ctx, card, tx.Amount)
	if err != nil {
		return fmt.Errorf("payment processing failed: %v", err)
	}

	return nil
}

func (s *PaymentServer) GetTransactionStatus(ctx context.Context, req *pb.GetTransactionStatusReq) (*pb.GetTransactionStatusResp, error) {
	tx, err := s.repo.GetTransactionByID(ctx, req.TransactionId)
	if err != nil {
		return nil, status.Error(codes.NotFound, "transaction not found")
	}

	return &pb.GetTransactionStatusResp{
		Status:  tx.Status,
		Amount:  float32(tx.Amount),
		OrderId: tx.OrderID,
	}, nil
}

func (s *PaymentServer) HealthCheck(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Status: true,
	}, nil
}
