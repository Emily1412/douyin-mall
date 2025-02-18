package service

import (
	"context"
	"errors"
	pb "user-service/api/user"
	"user-service/internal/model"
	"user-service/internal/pkg/jwt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginService struct {
	db  *gorm.DB
	ctx context.Context
}

func NewLoginService(ctx context.Context, db *gorm.DB) *LoginService {
	return &LoginService{
		ctx: ctx,
		db:  db,
	}
}

// Run 实现登录逻辑
func (s *LoginService) Run(req *pb.LoginReq) (resp *pb.LoginResp, err error) {
	// 查找用户
	userRow, err := model.GetByEmail(s.db, req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("用户不存在")
		}
		return nil, errors.New("查询用户失败")
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(userRow.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("密码错误")
	}

	// 生成 JWT token
	token, err := jwt.GenerateToken(userRow.ID)
	if err != nil {
		return nil, errors.New("生成token失败")
	}

	return &pb.LoginResp{
		UserId: int32(userRow.ID),
		Token:  token,
	}, nil
}
