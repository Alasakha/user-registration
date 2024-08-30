// models/user.go
package models

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

var JwtKey = []byte("secret_key")

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique"`
	Password string `json:"password"`
	Email    string `json:"email" gorm:"unique"`
	Role     string `json:"role"`
}

// Claims represents the JWT claims
type Claims struct {
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
