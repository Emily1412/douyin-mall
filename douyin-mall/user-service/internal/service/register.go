package service

import (
	"context"
	"errors"
	pb "user-service/api/user"
	"user-service/internal/model"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type RegisterService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewRegisterService(ctx context.Context, db *gorm.DB) *RegisterService {
	return &RegisterService{
		ctx: ctx,
		db:  db,
	}
}

// Run 实现注册逻辑
func (s *RegisterService) Run(req *pb.RegisterReq) (resp *pb.RegisterResp, err error) {
	// 验证密码
	if req.Password != req.ConfirmPassword {
		err = errors.New("两次输入的密码不一致")
		return
	}

	// 生成密码哈希
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("密码加密失败")
	}

	// 创建新用户
	newUser := &model.User{
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	// 保存到数据库
	if err = model.Create(s.db, newUser); err != nil {
		return nil, errors.New("创建用户失败")
	}

	return &pb.RegisterResp{
		UserId: int32(newUser.ID),
	}, nil
}
