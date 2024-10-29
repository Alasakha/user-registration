package models

import "time"

type Route struct {
	ID       int    `json:"id"`
	Path     string `json:"path"`
	Label    string `json:"label"`
	Icon     string `json:"icon"`
	ParentID *int   `json:"parent_id,omitempty"` // 父级路由
	IsActive bool   `json:"is_active"`
}

type MenuItem struct {
	ID       int         `json:"id"`
	Label    string      `json:"label"`
	Path     string      `json:"path"`
	Icon     string      `json:"icon"`
	Children []*MenuItem `json:"children,omitempty"` // 子菜单项
}

type Menu struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `json:"name"`
	Path      string    `json:"path"`
	ParentID  *uint     `json:"parent_id"`
	SortOrder int       `json:"sort_order"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Children  []Menu    `gorm:"-" json:"children"`
}
