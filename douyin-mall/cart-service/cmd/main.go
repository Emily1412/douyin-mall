package main

import (
	"douyin-mall/cart-service/config"
	"douyin-mall/cart-service/internal/handler"
	"douyin-mall/cart-service/internal/model"
	"douyin-mall/cart-service/internal/service"
	"douyin-mall/cart-service/proto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 加载配置
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		log.Fatalf("Load config error: %v", err)
	}

	// 初始化数据库
	config.InitDB()

	// 自动迁移表结构
	config.DB.AutoMigrate(&model.Cart{}, &model.CartItem{})
	fmt.Println("Database migrated")

	// 创建 CartService 实例
	cartService := service.NewCartService(config.DB)

	// 创建 CartHandler 实例
	cartHandler := handler.NewCartHandler(cartService)

	// 启动 gRPC 服务器
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	proto.RegisterCartServiceServer(grpcServer, cartHandler)

	fmt.Println("CartService is running on port 50053...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
