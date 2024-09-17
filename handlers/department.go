package handlers

import (
	"net/http"
	"user-registration/database"

	"github.com/gin-gonic/gin"
)

// 处理部门和岗位树形结构请求
func GetDepartmentTreeHandler(c *gin.Context) {
	// 获取部门和岗位的树形结构
	companies, err := database.GetDepartmentTree()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": companies,
		"code": 200,
	})
}
