// handlers/user.go
package handlers

import (
	"net/http"
	"user-registration/database"
	"user-registration/models"

	"github.com/gin-gonic/gin"
)

// 获取用户信息的处理函数
func GetUserInfo(c *gin.Context) {
	// 从上下文中获取用户名
	username := c.MustGet("username").(string)

	var user models.User
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	// 返回用户信息
	c.JSON(http.StatusOK, gin.H{
		"username": user.Username,
		"email":    user.Email,
		"role":     user.Role,
	})
}
