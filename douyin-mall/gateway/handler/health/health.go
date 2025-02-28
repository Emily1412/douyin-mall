package health

import (
	"context"
	"douyin-mall/gateway/config"

	//proto所在地
	pb "douyin-mall/payment-service/api"
	"net/http"
	"sync"
	"time"

	"google.golang.org/grpc"

	"github.com/gin-gonic/gin"
)

// 服务健康检查函数类型
type healthCheckFunc func() bool

// 服务健康检查映射
var serviceChecks = map[string]healthCheckFunc{
	"payment": checkPaymentService,
	// "user":     checkUserService,
	// "order":    checkOrderService,
	// "product":  checkProductService,
	// "cart":     checkCartService,
	// "checkout": checkCheckoutService,
}

func checkPaymentService() bool {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	conn, err := grpc.Dial(config.GlobalConfig.Services.Payment, grpc.WithInsecure())
	if err != nil {
		return false
	}
	defer conn.Close()

	client := pb.NewPaymentServiceClient(conn)
	resp, err := client.HealthCheck(ctx, &pb.HealthCheckRequest{})
	if err != nil {
		return false
	}
	return resp.Status
}

// 检查商品服务
// func checkProductService() bool {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	conn, err := grpc.Dial(config.GlobalConfig.Services.Product, grpc.WithInsecure())
// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()

// 	client := pb.NewProductServiceClient(conn)
// 	_, err = client.HealthCheck(ctx, &pb.HealthCheckRequest{})
// 	return err == nil
// }

// // 检查购物车服务
// func checkCartService() bool {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	conn, err := grpc.Dial(config.GlobalConfig.Services.Cart, grpc.WithInsecure())
// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()

// 	client := pb.NewCartServiceClient(conn)
// 	_, err = client.HealthCheck(ctx, &pb.HealthCheckRequest{})
// 	return err == nil
// }

// // 检查结算服务
// func checkCheckoutService() bool {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	conn, err := grpc.Dial(config.GlobalConfig.Services.Checkout, grpc.WithInsecure())
// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()

// 	client := pb.NewCheckoutServiceClient(conn)
// 	_, err = client.HealthCheck(ctx, &pb.HealthCheckRequest{})
// 	return err == nil
// }

// // 检查用户服务
// func checkUserService() bool {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	conn, err := grpc.Dial(config.GlobalConfig.Services.User, grpc.WithInsecure())
// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()

// 	return err == nil
// }

// // 检查订单服务
// func checkOrderService() bool {
// 	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
// 	defer cancel()

// 	conn, err := grpc.Dial(config.GlobalConfig.Services.Order, grpc.WithInsecure())
// 	if err != nil {
// 		return false
// 	}
// 	defer conn.Close()

// 	return err == nil
// }

// 并发检查所有服务健康状态
func checkAllServices() map[string]bool {
	results := make(map[string]bool)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for serviceName, checkFunc := range serviceChecks {
		wg.Add(1)
		go func(name string, check healthCheckFunc) {
			defer wg.Done()
			status := check()
			mutex.Lock()
			results[name] = status
			mutex.Unlock()
		}(serviceName, checkFunc)
	}

	wg.Wait()
	return results
}

// 健康检查处理函数
func HealthCheck(c *gin.Context) {
	services := checkAllServices()
	c.JSON(http.StatusOK, gin.H{
		"status":    "up",
		"services":  services,
		"timestamp": time.Now().Unix(),
	})
}

// 就绪检查
func ReadinessCheck(c *gin.Context) {
	services := checkAllServices()

	allReady := true
	for _, status := range services {
		if !status {
			allReady = false
			break
		}
	}

	if allReady {
		c.JSON(http.StatusOK, gin.H{
			"status":    "ready",
			"details":   services,
			"timestamp": time.Now().Unix(),
		})
	} else {
		c.JSON(http.StatusServiceUnavailable, gin.H{
			"status":    "not ready",
			"details":   services,
			"timestamp": time.Now().Unix(),
		})
	}
}
