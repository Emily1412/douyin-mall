// gRPC 拦截器
package middleware

import (
	"context"
	"douyin-mall/common/utils/logger"
	"douyin-mall/common/utils/metrics"
	"time"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// UnaryServerInterceptor 创建统一的 gRPC 拦截器
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()
		method := info.FullMethod

		// 记录请求开始
		logger.Log.Info("gRPC request started",
			zap.String("method", method))

		// 增加请求计数
		metrics.RPCRequests.WithLabelValues(method, "started").Inc()

		// 执行实际的处理方法
		resp, err := handler(ctx, req)

		// 记录处理时间
		duration := time.Since(start).Seconds()
		metrics.RPCDuration.WithLabelValues(method).Observe(duration)

		// 记录请求结果
		if err != nil {
			metrics.RPCRequests.WithLabelValues(method, "failed").Inc()
			logger.Log.Error("gRPC request failed",
				zap.String("method", method),
				zap.Error(err),
				zap.Duration("duration", time.Since(start)))
		} else {
			metrics.RPCRequests.WithLabelValues(method, "success").Inc()
			logger.Log.Info("gRPC request completed",
				zap.String("method", method),
				zap.Duration("duration", time.Since(start)))
		}

		return resp, err
	}
}
