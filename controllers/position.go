package controllers

import (
	"net/http"
	"user-registration/database"
	"user-registration/utils" // 引入通用响应的包

	"github.com/gin-gonic/gin"
)

type Position struct {
	ID            uint   `json:"id" gorm:"primary_key"`
	Name          string `json:"name"`
	Department_ID uint   `json:"department_id"`
}

// 获取所有职位信息
func GetPositions(c *gin.Context) {
	var positions []Position
	if err := database.DB.Find(&positions).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to retrieve positions", nil)
		return
	}
	utils.Respond(c, http.StatusOK, "Positions retrieved successfully", positions)
}

// 创建新职位
func CreatePosition(c *gin.Context) {
	var input Position
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}
	if err := database.DB.Create(&input).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to create position", nil)
		return
	}
	utils.Respond(c, http.StatusOK, "Position created successfully", input)
}

// 修改职位信息
func UpdatePosition(c *gin.Context) {
	id := c.Param("id")
	var position Position
	if err := database.DB.Where("id = ?", id).First(&position).Error; err != nil {
		utils.Respond(c, http.StatusNotFound, "Position not found", nil)
		return
	}

	var input Position
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Respond(c, http.StatusBadRequest, "Invalid input", nil)
		return
	}

	// 更新字段
	position.Name = input.Name
	position.Department_ID = input.Department_ID

	if err := database.DB.Save(&position).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to update position", nil)
		return
	}
	utils.Respond(c, http.StatusOK, "Position updated successfully", position)
}

// 删除职位
func DeletePosition(c *gin.Context) {
	id := c.Param("id")
	if err := database.DB.Delete(&Position{}, id).Error; err != nil {
		utils.Respond(c, http.StatusInternalServerError, "Failed to delete position", nil)
		return
	}
	utils.Respond(c, http.StatusOK, "Position deleted successfully", nil)
}
