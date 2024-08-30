// main.go
package main

import (
	"net/http"
	"user-registration/database"
	"user-registration/handlers"
	"user-registration/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 允许所有来源的CORS请求
	r.Use(cors.Default())

	// 初始化数据库
	database.Connect()

	// 设置根路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "API is running")
	})

	// 路由配置
	r.POST("/register", handlers.Register)

	// 登录路由
	r.POST("/login", handlers.Login)

	// 受保护的路由
	protected := r.Group("/home")
	protected.Use(middlewares.JWTAuthMiddleware())
	{
		protected.GET("/", func(c *gin.Context) {
			username := c.MustGet("username").(string)
			role := c.MustGet("role").(string)
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected route", "username": username, "role": role})
		})
	}

	// 启动服务
	r.Run(":8080")
}
