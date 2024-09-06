package models

type Route struct {
	ID       int    `json:"id"`
	Path     string `json:"path"`
	Label    string `json:"label"`
	Icon     string `json:"icon"`
	ParentID *int   `json:"parent_id,omitempty"` // 父级路由
	IsActive bool   `json:"is_active"`
}

type MenuItem struct {
	Label    string     `json:"label"`
	Path     string     `json:"path"`
	Icon     string     `json:"icon"`
	Children []MenuItem `json:"children,omitempty"` // 子菜单项
}
