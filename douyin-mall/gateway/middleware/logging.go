package middleware

import (
	"douyin-mall/common/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func LoggingMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path

		// 请求开始日志
		logger.Log.Info("API request started",
			zap.String("path", path),
			zap.String("method", c.Request.Method),
			zap.String("ip", c.ClientIP()))

		c.Next()

		// 请求结束日志
		logger.Log.Info("API request completed",
			zap.String("path", path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", time.Since(start)))
	}
}
