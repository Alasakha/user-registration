package handlers

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"user-registration/models"

	"user-registration/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// （1）创建菜单接口
func PostMenu(c *gin.Context) {
	var menu models.Menu
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := database.DB.Create(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法创建菜单"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "菜单创建成功", "data": menu})
}

// (2) 获取树形结构
func GetMenuList(c *gin.Context) {
	log.Println("Received request to get menu list") // 添加日志
	var menus []models.Menu
	if err := database.DB.Order("sort_order").Find(&menus).Error; err != nil {
		log.Printf("Database query error: %v", err) // 打印错误信息
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法获取菜单列表"})
		return
	}

	fmt.Printf("Fetched menus: %+v\n", menus) // 尝试用 fmt.Printf 打印
	// 构建树形结构的菜单列表
	menuTree := buildMenuTree(menus, nil)

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": menuTree,
	})
}

func buildMenuTree(menus []models.Menu, parentID *uint) []models.Menu {
	var tree []models.Menu
	for _, menu := range menus {
		if (parentID == nil && menu.ParentID == nil) || (parentID != nil && menu.ParentID != nil && *parentID == *menu.ParentID) {
			menu.Children = buildMenuTree(menus, &menu.ID)
			tree = append(tree, menu)
		}
	}
	return tree
}

// （3）更新菜单接口
func UpdateMenu(c *gin.Context) {
	var menu models.Menu
	id := c.Param("id") // 获取 ID

	// 根据 ID 查找菜单
	if err := database.DB.First(&menu, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "菜单未找到"})
		return
	}

	// 绑定 JSON 数据
	if err := c.ShouldBindJSON(&menu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Received menu data: %+v\n", menu) // 打印接收到的菜单数据

	// 更新菜单
	if err := database.DB.Save(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "无法更新菜单", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "菜单更新成功"})
}

// 通过 id 删除菜单
func DeleteMenu(c *gin.Context) {
	id := c.Param("id")

	// 检查菜单是否存在
	var menu models.Menu
	if err := database.DB.First(&menu, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "菜单未找到"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "查询菜单失败"})
		return
	}

	// 删除菜单
	if err := database.DB.Delete(&menu).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "删除失败", "details": err.Error()})
		return
	}

	// 记录日志（可选）
	log.Printf("菜单 ID %s 已被删除", id)

	c.JSON(http.StatusOK, gin.H{"message": "删除成功", "deleted_id": id})
}
