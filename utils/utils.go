package utils

import (
	"github.com/gin-gonic/gin"
)

// 通用响应函数
func Respond(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"data":    data,
	})
}
