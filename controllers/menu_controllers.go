package controllers

import (
	"net/http"
	"user-registration/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetMenu(c *gin.Context) {
	roleID := 1 // 假设从JWT中解析出的用户角色ID

	db := c.MustGet("db").(*gorm.DB) // 从上下文中获取数据库连接
	menu, err := services.GetMenuByRole(roleID, db)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load menu"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": menu,
	})
}
