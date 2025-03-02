package main

import (
	"fmt"
	"log"
	"net"

	"douyin-mall/payment-service/config"
	"douyin-mall/payment-service/internal/handler"
	"douyin-mall/payment-service/internal/repository"
	"douyin-mall/payment-service/internal/service"
	pb "douyin-mall/payment-service/proto" // 确保这里是正确的包路径

	"google.golang.org/grpc"
)

func main() {
	// 加载配置文件
	if err := config.LoadConfig(); err != nil {
		log.Fatalf("Load config error: %v", err)
	}

	// 初始化数据库
	config.InitDB()

	// 创建 PaymentRepository 实例
	paymentRepository := repository.NewPaymentRepository(config.DB)

	// 创建 PaymentService 实例
	paymentService := service.NewPaymentService(paymentRepository)

	// 创建 PaymentHandler 实例
	paymentHandler := handler.NewPaymentHandler(paymentService)

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterPaymentServiceServer(grpcServer, paymentHandler) // 注册服务

	fmt.Println("PaymentService is running on port 50055...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
