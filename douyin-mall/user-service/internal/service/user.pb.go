package service

import (
	"context"
	pb "user-service/api/user"

	"gorm.io/gorm"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	DB *gorm.DB
}

// Register 实现用户注册
func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	registerService := NewRegisterService(ctx, s.DB)
	return registerService.Run(req)
}

// Login 实现用户登录
func (s *UserService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginResp, error) {
	loginService := NewLoginService(ctx, s.DB)
	return loginService.Run(req)
}
