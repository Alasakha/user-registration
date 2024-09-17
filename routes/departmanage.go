package routes

import (
	"user-registration/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterPositionRoutes(r *gin.Engine) {
	positionGroup := r.Group("/positions")
	{
		positionGroup.GET("/", controllers.GetPositions)         // 获取职位列表
		positionGroup.POST("/", controllers.CreatePosition)      // 创建职位
		positionGroup.PUT("/:id", controllers.UpdatePosition)    // 修改职位
		positionGroup.DELETE("/:id", controllers.DeletePosition) // 删除职位
	}
}
