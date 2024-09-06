// middleware/auth_middleware.go
package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 模拟从 JWT 解析出用户角色 ID
		roleID := 1 // 实际应该从 JWT token 中获取
		c.Set("role_id", roleID)

		c.Next()
	}
}
