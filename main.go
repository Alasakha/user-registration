package main

import (
	"log"
	"net/http"
	"user-registration/controllers"
	"user-registration/database" // 确保正确引入
	"user-registration/handlers"
	"user-registration/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 配置CORS中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // 前端应用的URL
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// 初始化数据库
	err := database.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer func() {
		sqlDB, _ := database.DB.DB()
		sqlDB.Close()
	}()

	// 将数据库连接添加到 Gin 的上下文
	r.Use(func(c *gin.Context) {
		c.Set("db", database.DB)
		c.Next()
	})

	// 设置根路由
	r.GET("/", func(c *gin.Context) {
		c.String(200, "API is running")
	})

	// 注册路由
	// r.POST("/register", handlers.Register)

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
		protected.GET("/user/info", handlers.GetUserInfo) // 获取用户信息的路由

		// 新增的菜单路由，基于角色从数据库返回动态菜单
		protected.GET("/menu", controllers.GetMenu)

		// 注册路由
		protected.POST("/register", handlers.Register)
	}

	// 启动服务
	r.Run(":8080")
}
