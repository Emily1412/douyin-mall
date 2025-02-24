package router

import (
	"douyin-mall/gateway/middleware"
	"douyin-mall/gateway/router/routes"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"golang.org/x/time/rate"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.CORSMiddleware())
	r.Use(middleware.LoggingMiddleware())

	// 限流器
	limiter := middleware.NewIPRateLimiter(rate.Limit(1), 5)
	r.Use(middleware.RateLimitMiddleware(limiter))

	// 监控指标
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))

	// 健康检查
	routes.RegisterHealthRoutes(r)

	// API 路由组
	api := r.Group("/api/v1")
	{
		routes.RegisterPaymentRoutes(api)
		// 后续添加其他服务路由
		// routes.RegisterUserRoutes(api)
        // routes.RegisterProductRoutes(api)
        // routes.RegisterCartRoutes(api)
        // routes.RegisterOrderRoutes(api)
        // routes.RegisterCheckoutRoutes(api)
	}

	return r
}
