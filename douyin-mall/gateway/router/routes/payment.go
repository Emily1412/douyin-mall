package routes

import (
	"douyin-mall/gateway/handler/payment"
	"douyin-mall/gateway/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterPaymentRoutes(r *gin.RouterGroup) {
	payments := r.Group("/payments")
	payments.Use(middleware.AuthMiddleware())
	{
		payments.POST("/charge", payment.Charge)
		payments.GET("/:id", payment.Status)
	}
}
