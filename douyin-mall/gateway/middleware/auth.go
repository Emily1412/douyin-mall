package middleware

import (
	"douyin-mall/common/utils/jwt"
	"encoding/csv"
	"net/http"
	"os"
	"strings"

	"log"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization token required"})
			c.Abort()
			return
		}

		// 去掉 Bearer 前缀
		token = strings.TrimPrefix(token, "Bearer ")

		claims, err := jwt.ParseToken(token)
		if err != nil {
			log.Printf("Token validation failed: %v", err)
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// 将用户信息存储到上下文中
		c.Set("user_id", claims.UserID)

		id := c.Param("id")
		if !hasAccess(claims.UserID, "/api/payments/"+id) {
			c.JSON(http.StatusForbidden, gin.H{"error": "permission denied"})
			c.Abort()
			return
		}

		c.Next()
	}
}

// hasAccess checks if the user has access to the specified path
func hasAccess(userID uint, path string) bool {
	// 读取 policy.csv 文件
	file, err := os.Open("douyin-mall/user-service/configs/policy.csv")
	if err != nil {
		return false // 文件打开失败，返回无权限
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return false // 读取文件失败，返回无权限
	}

	// 遍历权限记录
	for _, record := range records {
		if len(record) < 3 {
			continue // 跳过无效记录
		}
		permission := record[1]
		apiPath := record[2]

		// 检查用户角色和请求路径
		if strings.Contains(apiPath, path) && permission == "access" {
			return true // 找到匹配的权限，返回有权限
		}
	}

	return false // 没有找到匹配的权限，返回无权限
}
