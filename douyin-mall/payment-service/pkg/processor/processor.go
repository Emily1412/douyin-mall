// 支付处理器
package payment

import (
	"context"
	//proto所在地
	pb "payment-service/api"
	"fmt"
	"math/rand"
	"time"
)

// PaymentProcessor 支付处理器接口
type PaymentProcessor interface {
	Process(ctx context.Context, cardInfo *pb.CreditCardInfo, amount float64) error
	Validate(cardInfo *pb.CreditCardInfo) error
}

// CreditCardProcessor 信用卡支付处理器
type CreditCardProcessor struct {
	// 可以添加配置，如商户ID、密钥等
	merchantID string
	secretKey  string
}

func NewCreditCardProcessor() *CreditCardProcessor {
	return &CreditCardProcessor{
		merchantID: "TEST_MERCHANT",
		secretKey:  "TEST_SECRET",
	}
}

// Validate 验证信用卡信息
func (p *CreditCardProcessor) Validate(cardInfo *pb.CreditCardInfo) error {
	// 1. 检查卡号长度（通常是13-19位）
	if len(cardInfo.CreditCardNumber) < 13 || len(cardInfo.CreditCardNumber) > 19 {
		return fmt.Errorf("invalid card number length")
	}

	// 2. 检查CVV（通常是3-4位）
	if cardInfo.CreditCardCvv < 100 || cardInfo.CreditCardCvv > 9999 {
		return fmt.Errorf("invalid CVV")
	}

	// 3. 检查过期日期
	currentYear := time.Now().Year()
	if cardInfo.CreditCardExpirationYear < int32(currentYear) {
		return fmt.Errorf("card has expired")
	}
	if cardInfo.CreditCardExpirationYear == int32(currentYear) &&
		cardInfo.CreditCardExpirationMonth < int32(time.Now().Month()) {
		return fmt.Errorf("card has expired")
	}

	return nil
}

// Process 处理支付
func (p *CreditCardProcessor) Process(ctx context.Context, cardInfo *pb.CreditCardInfo, amount float64) error {
	// 1. 验证信用卡信息
	if err := p.Validate(cardInfo); err != nil {
		return fmt.Errorf("card validation failed: %v", err)
	}

	// 2. 模拟风控检查
	if err := p.riskCheck(amount); err != nil {
		return fmt.Errorf("risk check failed: %v", err)
	}

	// 3. 模拟支付网关处理
	if err := p.mockGatewayProcess(ctx); err != nil {
		return fmt.Errorf("payment gateway error: %v", err)
	}

	return nil
}

// riskCheck 模拟风控检查
func (p *CreditCardProcessor) riskCheck(amount float64) error {
	// 模拟高额交易检查
	if amount > 50000 {
		return fmt.Errorf("amount exceeds maximum limit")
	}

	// 模拟随机风控拦截（1%的概率）
	if rand.Float64() < 0.01 {
		return fmt.Errorf("transaction blocked by risk control")
	}

	return nil
}

// mockGatewayProcess 模拟支付网关处理
func (p *CreditCardProcessor) mockGatewayProcess(ctx context.Context) error {
	// 模拟处理时间
	select {
	case <-ctx.Done():
		return fmt.Errorf("payment processing timeout")
	case <-time.After(time.Duration(rand.Int63n(2000)) * time.Millisecond):
		// 模拟支付失败（5%的概率）
		if rand.Float64() < 0.05 {
			return fmt.Errorf("payment declined by bank")
		}
		return nil
	}
}
