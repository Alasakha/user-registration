// handlers/user.go
package handlers

import (
	"net/http"
	"os"
	"time"
	"user-registration/database"
	"user-registration/models"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

// JWT密钥应从环境变量中读取，增强安全性
var JwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// Claims 结构，用于生成 JWT
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.RegisteredClaims
}

func Login(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "Invalid input",
			"error":   err.Error(),
		})
		return
	}

	var user models.User
	// 检查用户名是否存在
	if err := database.DB.Where("username = ?", input.Username).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "用户名不存在",
		})
		return
	}

	// 验证密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"code":    401,
			"message": "密码或账号输入错误",
		})
		return
	}

	// 生成JWT Token
	expirationTime := time.Now().Add(time.Hour * 1) // Token 1小时后过期
	claims := &Claims{
		Username: user.Username,
		Role:     user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JwtKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "Failed to generate token",
			"error":   err.Error(),
		})
		return
	}

	// 返回token和成功信息
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "Login successful",
		"token":   tokenString,
	})
}
