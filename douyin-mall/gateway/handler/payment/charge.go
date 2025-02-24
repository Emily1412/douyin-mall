package payment

import (
	"context"
	"douyin-mall/gateway/config"
	//proto所在地
	pb "douyin-mall/payment-service/api"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func Charge(c *gin.Context) {
	// 原来 PaymentChargeHandler 的实现
	// 从上下文获取用户ID
	userID, _ := c.Get("userId")

	// 解析请求
	var req struct {
		Amount     float32 `json:"amount"`
		OrderID    string  `json:"order_id"`
		CreditCard struct {
			Number          string `json:"number"`
			CVV             int32  `json:"cvv"`
			ExpirationYear  int32  `json:"expiration_year"`
			ExpirationMonth int32  `json:"expiration_month"`
		} `json:"credit_card"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用支付服务
	conn, err := grpc.Dial(config.GlobalConfig.Services.Payment, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "service unavailable"})
		return
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewPaymentServiceClient(conn)

	// 调用远程服务
	resp, err := client.Charge(context.Background(), &pb.ChargeReq{
		Amount:  req.Amount,
		OrderId: req.OrderID,
		UserId:  userID.(uint32),
		CreditCard: &pb.CreditCardInfo{
			CreditCardNumber:          req.CreditCard.Number,
			CreditCardCvv:             req.CreditCard.CVV,
			CreditCardExpirationYear:  req.CreditCard.ExpirationYear,
			CreditCardExpirationMonth: req.CreditCard.ExpirationMonth,
		},
	})

	// 处理响应	
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transaction_id": resp.TransactionId})
}

func Status(c *gin.Context) {
	// 原来 PaymentStatusHandler 的实现
	transactionID := c.Param("id")

	// 调用支付服务
	conn, err := grpc.Dial(config.GlobalConfig.Services.Payment, grpc.WithInsecure())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "service unavailable"})
		return
	}
	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)
	resp, err := client.GetTransactionStatus(context.Background(), &pb.GetTransactionStatusReq{
		TransactionId: transactionID,
	})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":   resp.Status,
		"amount":   resp.Amount,
		"order_id": resp.OrderId,
	})
}
