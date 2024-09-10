package controllers

import (
	"net/http"
	"user-registration/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// func GetMenu(c *gin.Context) {
// 	role, exists := c.Get("role")
// 	if !exists {
// 		c.JSON(http.StatusUnauthorized, gin.H{"error": "Role not found in context"})
// 		return
// 	}
// 	roleStr, ok := role.(string)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid role type"})
// 		return
// 	}

// 	db, exists := c.Get("db")
// 	if !exists {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection not found in context"})
// 		return
// 	}
// 	dbConn, ok := db.(*gorm.DB)
// 	if !ok {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid database connection"})
// 		return
// 	}

// 	menu, err := services.GetMenuByRole(roleStr, dbConn)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load menu", "details": err.Error()})
// 		return
// 	}

//		c.JSON(http.StatusOK, gin.H{
//			"code": 200,
//			"data": menu,
//		})
//	}
//
// GetMenu 根据用户角色返回菜单
func GetMenu(c *gin.Context) {
	role := c.MustGet("role").(string) // 从 JWT 中间件获取用户角色

	// 从上下文获取数据库连接
	db := c.MustGet("db").(*gorm.DB)

	// 调用服务获取菜单
	menuItems, err := services.GetMenuByRole(role, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch menu"})
		return
	}

	// 返回菜单数据
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": menuItems,
	})
}
