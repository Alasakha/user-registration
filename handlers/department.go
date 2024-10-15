package handlers

import (
	"net/http"
	"user-registration/database"
	"user-registration/models"

	"github.com/gin-gonic/gin"
)

// RequestBody 定义了创建部门时的请求体结构
type RequestBody struct {
	Name      string `json:"name" binding:"required"`
	CompanyID uint   `json:"company_id" binding:"required"`
	ParentID  *uint  `json:"parent_id"`
	Status    string `json:"status"`
	SortOrder int    `json:"sort_order"`
}

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

// 创建部门
func PostDepartment(c *gin.Context) {
	var reqBody RequestBody

	// 绑定传入的 JSON 到 RequestBody 结构体
	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "无法创建部门",
			"details": err.Error(),
		})
		return
	}

	//创建新的部门
	department := models.Department{
		Name:      reqBody.Name,
		CompanyID: reqBody.CompanyID,
		ParentID:  reqBody.ParentID,
		Status:    reqBody.Status,
		SortOrder: reqBody.SortOrder,
	}

	// 保存部门到数据库
	if err := database.DB.Create(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建部门"})
		return
	}

	// 返回创建的部门信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    department,
		"message": "部门信息上传成功",
	})
}
