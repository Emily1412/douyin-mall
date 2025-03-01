package main

import (
	"douyin-mall/gateway/config"
	"douyin-mall/gateway/router"
	"douyin-mall/common/utils/logger"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	if err := config.Init(); err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}

	// 初始化日志
	if err := logger.Init(); err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.Log.Sync()

	// 设置 Gin 模式
	if config.GlobalConfig.Server.Mode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// 设置路由
	r := router.SetupRouter()

	// 记录服务启动日志
	addr := fmt.Sprintf(":%d", config.GlobalConfig.Server.Port)
	logger.Log.Info("Gateway server starting",
		zap.String("address", addr),
		zap.String("mode", config.GlobalConfig.Server.Mode))

	// 启动服务器
	if err := r.Run(addr); err != nil {
		logger.Log.Fatal("Failed to start server",
			zap.Error(err))
	}
}
