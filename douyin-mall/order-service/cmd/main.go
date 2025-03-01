package main

import (
	"fmt"
	"log"
	"net"

	"douyin-mall/order-service/config"
	"douyin-mall/order-service/internal/handler"
	"douyin-mall/order-service/internal/model"
	"douyin-mall/order-service/internal/service"
	pb "douyin-mall/order-service/proto"

	"google.golang.org/grpc"
)

func main() {
	// 加载配置文件
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		log.Fatalf("Load config error: %v", err)
	}

	// 初始化数据库
	config.InitDB()

	// 自动迁移订单相关模型
	config.DB.AutoMigrate(&model.Order{}, &model.OrderItem{})
	fmt.Println("Database migrated")

	// 创建 OrderService 实例
	orderService := service.NewOrderService(config.DB)

	// 创建 OrderHandler 实例
	orderHandler := handler.NewOrderHandler(orderService)

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, orderHandler)

	fmt.Println("OrderService is running on port 50054...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
