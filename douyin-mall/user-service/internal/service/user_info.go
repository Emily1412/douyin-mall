package service

import (
	"context"
	"errors"
	"fmt"
	pb "user-service/api"
	"user-service/internal/model"

	"gorm.io/gorm"
)

type UserInfoService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewUserInfoService(ctx context.Context, db *gorm.DB) *UserInfoService {
	return &UserInfoService{
		ctx: ctx,
		db:  db,
	}
}

// Run 实现获取用户信息逻辑
func (s *UserInfoService) Run(req *pb.GetUserInfoReq) (resp *pb.GetUserInfoResp, err error) {
	fmt.Printf("Received get user info request: %+v\n", req)

	// 查找用户
	var user model.User
	err = s.db.First(&user, req.UserId).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, errors.New("查询用户失败")
	}

	return &pb.GetUserInfoResp{
		UserId:    int32(user.ID),
		Email:     user.Email,
		CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
	}, nil
} 