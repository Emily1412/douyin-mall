package middleware

import (
	"context"
	"fmt"
	"path/filepath"
	"user-service/internal/pkg/jwt"

	"github.com/casbin/casbin/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type AuthMiddleware struct {
	enforcer *casbin.Enforcer
}

func NewAuthMiddleware() (*AuthMiddleware, error) {
	// 获取当前工作目录的上级目录（项目根目录）
	currentDir, err := filepath.Abs(".")
	if err != nil {
		return nil, err
	}
	rootDir := filepath.Dir(currentDir)

	// 构建配置文件的完整路径
	modelPath := filepath.Join(rootDir, "configs", "rbac_model.conf")
	policyPath := filepath.Join(rootDir, "configs", "policy.csv")

	// 初始化 CasBin
	enforcer, err := casbin.NewEnforcer(modelPath, policyPath)
	if err != nil {
		return nil, err
	}

	return &AuthMiddleware{enforcer: enforcer}, nil
}

func (m *AuthMiddleware) AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 跳过登录和注册接口的验证
	if info.FullMethod == "/user.UserService/Login" || info.FullMethod == "/user.UserService/Register" {
		return handler(ctx, req)
	}

	// 从 metadata 中获取 token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "无法获取 metadata")
	}

	tokens := md.Get("authorization")
	if len(tokens) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "未提供 token")
	}

	// 验证 token
	claims, err := jwt.ParseToken(tokens[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "无效的 token")
	}

	// CasBin 权限验证
	fmt.Printf("Checking permission for user %d, method %s\n", claims.UserID, info.FullMethod)
	// 将 uint 类型转换为字符串
	userID := fmt.Sprintf("%d", claims.UserID)
	allowed, err := m.enforcer.Enforce(userID, info.FullMethod, "access")
	if err != nil {
		fmt.Printf("Permission check error: %v\n", err)
		return nil, status.Errorf(codes.Internal, "权限验证失败")
	}
	if !allowed {
		fmt.Printf("Permission denied for user %d\n", claims.UserID)
		return nil, status.Errorf(codes.PermissionDenied, "没有访问权限")
	}

	// 将用户信息添加到 context
	newCtx := context.WithValue(ctx, "user_id", claims.UserID)
	return handler(newCtx, req)
}
