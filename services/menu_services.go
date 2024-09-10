package services

import (
	"fmt"
	"user-registration/models"

	"gorm.io/gorm"
)

// GetMenuByRole 根据角色从数据库获取层级菜单
func GetMenuByRole(role string, db *gorm.DB) ([]*models.MenuItem, error) {
	// 查询所有角色对应的菜单项，包括一级和二级菜单
	query := `
		SELECT r.id, r.path, r.label, r.icon, r.parent_id
		FROM routes r
		JOIN role_routes rr ON r.id = rr.route_id
		WHERE rr.role = ?;
	`
	rows, err := db.Raw(query, role).Rows()
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	menuMap := make(map[int]*models.MenuItem)
	var rootMenuItems []*models.MenuItem

	for rows.Next() {
		var id, parentID *int
		var path, label, icon string

		if err := rows.Scan(&id, &path, &label, &icon, &parentID); err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}

		fmt.Printf("Processing menu item: id=%d, parentID=%v, path=%s, label=%s\n", *id, parentID, path, label)

		menuItem := &models.MenuItem{
			ID:    *id,
			Label: label,
			Path:  path,
			Icon:  icon,
		}

		menuMap[*id] = menuItem

		if parentID == nil {
			rootMenuItems = append(rootMenuItems, menuItem)
		} else {
			if parent, exists := menuMap[*parentID]; exists {
				if parent.Children == nil {
					parent.Children = []*models.MenuItem{}
				}
				parent.Children = append(parent.Children, menuItem)
				fmt.Printf("Added child menu item: %+v to parent: %+v\n", *menuItem, *parent)
			} else {
				fmt.Printf("Parent menu item with ID %d not found\n", *parentID)
			}
		}
	}

	// Check for any errors that may have occurred during iteration
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	// Log for debugging
	fmt.Printf("Menu Map: %+v\n", menuMap)
	fmt.Printf("\n")
	fmt.Printf("Root Menu Items: %+v\n", rootMenuItems)

	return rootMenuItems, nil
}
