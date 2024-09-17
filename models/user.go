// models/user.go
package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username      string `json:"username" gorm:"unique"`
	Password      string `json:"password"`
	Email         string `json:"email" gorm:"unique"`
	Role          string `json:"role"`
	Department_id string `json:"department_id"`
	Position_id   string `json:"position_id"`
	Department    string `json:"department"`
	Position      string `json:"position"`
}
