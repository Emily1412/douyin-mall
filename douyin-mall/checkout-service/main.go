package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"checkout-service/config"
	"checkout-service/internal/service"
	pb "checkout-service/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// 设置日志格式
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 创建 gRPC 服务器
	server := grpc.NewServer()

	// 注册结算服务
	checkoutService := service.NewCheckoutService(cfg)
	pb.RegisterCheckoutServiceServer(server, checkoutService)

	// 注册反射服务，用于 grpcurl 等工具
	reflection.Register(server)

	// 监听端口
	addr := fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("监听失败: %v", err)
	}

	// 启动服务
	log.Printf("结算服务启动在 %s...", addr)

	// 在 goroutine 中启动服务
	go func() {
		if err := server.Serve(lis); err != nil {
			log.Fatalf("服务启动失败: %v", err)
		}
	}()

	// 等待中断信号优雅地关闭服务器
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("正在关闭结算服务...")
	server.GracefulStop()
	log.Println("结算服务已停止")
}
