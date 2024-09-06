package services

import (
	"database/sql"
	"user-registration/models"

	"gorm.io/gorm"
)

// 从数据库根据角色ID获取菜单
func GetMenuByRole(roleID int, db *gorm.DB) ([]models.MenuItem, error) {
	query := `
    SELECT r.id, r.path, r.label, r.icon, r.parent_id
    FROM routes r
    JOIN role_routes rr ON r.id = rr.route_id
    WHERE rr.role_id = ?
    ORDER BY r.parent_id, r.id
    `
	// 使用 Gorm 的 Raw 方法执行原始 SQL 查询
	rows, err := db.Raw(query, roleID).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var menuItems []models.MenuItem
	var menuMap = make(map[int]*models.MenuItem)

	for rows.Next() {
		var id, parentID sql.NullInt64
		var path, label, icon sql.NullString

		if err := rows.Scan(&id, &path, &label, &icon, &parentID); err != nil {
			return nil, err
		}

		menuItem := models.MenuItem{
			Label: label.String,
			Path:  path.String,
			Icon:  icon.String,
		}

		if parentID.Valid {
			// 子菜单
			if parentItem, ok := menuMap[int(parentID.Int64)]; ok {
				parentItem.Children = append(parentItem.Children, menuItem)
			}
		} else {
			// 顶级菜单
			menuItems = append(menuItems, menuItem)
			menuMap[int(id.Int64)] = &menuItem
		}
	}

	return menuItems, nil
}
