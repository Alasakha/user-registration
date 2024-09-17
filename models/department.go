package models

// // 树节点的结构
// type TreeNode struct {
// 	ID       uint       `json:"id"`
// 	Label    string     `json:"label"`
// 	Children []TreeNode `json:"children,omitempty"`
// }

// type Company struct {
// 	ID       uint         `gorm:"primaryKey"`
// 	Name     string       `gorm:"size:255"`
// 	Children []Department `gorm:"foreignKey:CompanyID"` // 添加Children字段
// }

// type Department struct {
// 	ID        uint         `gorm:"primaryKey"`
// 	Name      string       `gorm:"size:255"`
// 	CompanyID uint         `gorm:"not null"`
// 	ParentID  *uint        `gorm:"default:null"`
// 	Children  []Department `gorm:"foreignKey:ParentID"`
// 	Positions []Position   `gorm:"foreignKey:DepartmentID"` // 改为实际外键
// }

// type Position struct {
// 	ID           uint   `gorm:"primaryKey"`
// 	Name         string `gorm:"size:255"`
// 	DepartmentID uint   `gorm:"not null"`
// }

// 树节点的结构
type TreeNode struct {
	ID       uint       `json:"id"`
	Label    string     `json:"label"`
	Children []TreeNode `json:"children,omitempty"`
}

type Company struct {
	ID        uint         `gorm:"primaryKey"`
	Name      string       `gorm:"size:255"`
	CreatedAt string       `gorm:"autoCreateTime"`
	Children  []Department `gorm:"foreignKey:CompanyID"` // 添加Children字段
}

type Department struct {
	ID        uint         `gorm:"primaryKey"`
	Name      string       `gorm:"size:255"`
	CompanyID uint         `gorm:"not null"`
	ParentID  *uint        `gorm:"default:null"`
	Status    string       `gorm:"size:10;default:'enabled'"`
	SortOrder int          `gorm:"default:0"`
	CreatedAt string       `gorm:"autoCreateTime"`
	Children  []Department `gorm:"foreignKey:ParentID"`
	Positions []Position   `gorm:"foreignKey:DepartmentID"` // 改为实际外键
}

type Position struct {
	ID           uint   `gorm:"primaryKey"`
	Name         string `gorm:"size:255"`
	DepartmentID uint   `gorm:"not null"`
	Status       string `gorm:"size:10;default:'enabled'"`
	SortOrder    int    `gorm:"default:0"`
	CreatedAt    string `gorm:"autoCreateTime"`
}
