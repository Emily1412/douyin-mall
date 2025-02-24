package routes

import (
	"douyin-mall/gateway/handler/health"

	"github.com/gin-gonic/gin"
)

func RegisterHealthRoutes(r *gin.Engine) {
	r.GET("/health", health.HealthCheck)
	r.GET("/ready", health.ReadinessCheck)
}
