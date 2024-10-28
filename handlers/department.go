package handlers

import (
	"log"
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

func DeleteDepartment(c *gin.Context) {

	role, exists := c.Get("role")
	if !exists || role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "没有获得权限",
			"message": "没有获得权限",
		})
		return
	}

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error":   "权限不足，只有管理员可以执行",
			"message": "权限不足",
		})
		return
	}
	id := c.Param("id")

	// 在数据库中查找该部门
	var department models.Department
	if err := database.DB.Where("id = ?", id).First(&department).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "部门不存在",
			"message": "err.Error()",
		})
		return
	}

	// 删除该部门
	if err := database.DB.Delete(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "无法删除部门",
			"message": err.Error(),
		})
		return
	}

	// 成功删除后，返回成功信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "部门删除成功",
	})
}
func PutDepartment(c *gin.Context) {
	// 定义结构体用于绑定 URI 参数
	var uriParams struct {
		ID uint `uri:"id" binding:"required"`
	}

	// 绑定 URI 参数
	if err := c.ShouldBindUri(&uriParams); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid department ID"})
		return
	}

	log.Printf("Department ID: %d", uriParams.ID) // 确认ID

	// 从请求体中获取要更新的数据
	var input struct {
		Name string `json:"name" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 查找部门
	var department models.Department
	if err := database.DB.First(&department, uriParams.ID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "查找不到部门"})
		return
	}

	// 更新部门名称并保存
	department.Name = input.Name
	if err := database.DB.Save(&department).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update department"})
		return
	}

	// 返回更新后的部门信息
	c.JSON(http.StatusOK, gin.H{"data": department})
}
