package main

import (
	"fmt"
	"log"
	"net"

	pb "user-service/api/user"
	"user-service/internal/middleware"
	"user-service/internal/service"

	"google.golang.org/grpc"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	dsn := "root:123456@tcp(localhost:3306)/bytedance_user?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// 创建 gRPC 服务器，添加拦截器
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.AuthInterceptor),
	)
	userService := &service.UserService{
		DB: db,
	}
	pb.RegisterUserServiceServer(server, userService)

	fmt.Println("Server listening at :50051")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
