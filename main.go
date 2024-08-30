// main.go
package main

import (
	"net/http"
	"user-registration/database"
	"user-registration/handlers"

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

	authorized := r.Group("/home")
	authorized.Use(handlers.AuthMiddleware())
	{
		authorized.GET("/", func(c *gin.Context) {
			role, _ := c.Get("role")
			c.JSON(http.StatusOK, gin.H{"message": "Welcome to the home page!", "role": role})
		})
	}
	// 启动服务
	r.Run(":8080")
}
