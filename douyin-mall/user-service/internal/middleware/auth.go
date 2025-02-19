package middleware

import (
	"context"
	"user-service/internal/pkg/jwt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	// 跳过登录和注册接口的验证
	if info.FullMethod == "/user.UserService/Login" || info.FullMethod == "/user.UserService/Register" {
		return handler(ctx, req)
	}

	// 从 metadata 中获取 token
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Unauthenticated, "无法获取 metadata")
	}

	token := md.Get("authorization")
	if len(token) == 0 {
		return nil, status.Errorf(codes.Unauthenticated, "未提供 token")
	}

	// 验证 token
	claims, err := jwt.ParseToken(token[0])
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "无效的 token")
	}

	// 将用户信息添加到 context
	newCtx := context.WithValue(ctx, "user_id", claims.UserID)
	return handler(newCtx, req)
}
