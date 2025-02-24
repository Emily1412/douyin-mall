package main

import (
	"douyin-mall/payment-service/internal/service"
	"douyin-mall/pkg/db"
	"douyin-mall/common/utils/logger"
	"douyin-mall/common/middleware"
	//proto所在地
	pb "douyin-mall/payment-service/api"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 创建监听器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// 初始化数据库连接
	config := db.Config{
		Username: "root", // 你的数据库用户名
		Password: "root", // 你的数据库密码
		Host:     "127.0.0.1",
		Port:     3306,
		DBName:   "douyin-mall", // 你的数据库名
	}

	if err := db.InitWithConfig(config); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 在初始化数据库连接后添加
	if err := db.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 测试数据库连接
	if err := db.TestConnection(); err != nil {
		log.Fatalf("Database connection test failed: %v", err)
	}
	log.Println("Successfully connected to database")

	// 创建 gRPC 服务器，添加拦截器
	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryServerInterceptor()),
	)

	// 注册服务
	pb.RegisterPaymentServiceServer(s, service.NewPaymentServer())

	log.Printf("Payment service is running on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
