// handlers/pagecut.go
package handlers

import (
	"net/http"
	"strconv"
	"user-registration/database" // 替换为你的模块路径

	"github.com/gin-gonic/gin"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
	Role      string `json:"role"`
}

// Pagecut 实现分页接口
func Pagecut(c *gin.Context) {
	// 获取分页参数
	pageStr := c.Query("page")
	pageSizeStr := c.Query("pageSize")

	// 默认值处理
	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize <= 0 {
		pageSize = 10
	}

	// 查询数据库
	var users []User
	var total int64

	// 获取数据库实例
	db := database.DB

	// 查询用户总数
	db.Model(&User{}).Count(&total)

	// 计算偏移量，查询分页数据
	offset := (page - 1) * pageSize
	db.Offset(offset).Limit(pageSize).Find(&users)

	// 返回分页数据
	c.JSON(http.StatusOK, gin.H{
		"code":     200,
		"data":     users,
		"total":    total,
		"page":     page,
		"pageSize": pageSize,
	})
}
