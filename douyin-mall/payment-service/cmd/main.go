package main

import (
	"douyin-mall/common/middleware"
	"douyin-mall/common/utils/logger"
	"payment-service/config"
	"payment-service/internal/service"
	"payment-service/pkg/database"

	//proto所在地
	pb "payment-service/api"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// 创建监听器
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.GlobalConfig.Server.Port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	// 初始化数据库连接
	dbConfig := database.Config{
		Host:     config.GlobalConfig.Database.Host,
		Port:     config.GlobalConfig.Database.Port,
		Username: config.GlobalConfig.Database.Username,
		Password: config.GlobalConfig.Database.Password,
		DBName:   config.GlobalConfig.Database.DBName,
	}

	if err := database.InitWithConfig(dbConfig); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// 在初始化数据库连接后添加
	if err := database.AutoMigrate(); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// 测试数据库连接
	if err := database.TestConnection(); err != nil {
		log.Fatalf("Database connection test failed: %v", err)
	}
	log.Println("Successfully connected to database")

	// 创建 gRPC 服务器，添加拦截器
	s := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.UnaryServerInterceptor()),
	)

	// 注册服务
	pb.RegisterPaymentServiceServer(s, service.NewPaymentServer())

	log.Printf("Payment service is running on :%d", config.GlobalConfig.Server.Port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
