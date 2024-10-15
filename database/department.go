package database

import (
	"user-registration/models"
)

// 获取部门和岗位的树形结构
func GetDepartmentTree() ([]models.Company, error) {
	var companies []models.Company
	var departments []models.Department
	var positions []models.Position

	// 查询所有公司
	if err := DB.Preload("Children").Find(&companies).Error; err != nil {
		return nil, err
	}

	// 查询所有活跃部门
	if err := DB.Where("status = ?", "enabled").Find(&departments).Error; err != nil {
		return nil, err
	}

	// 查询所有活跃岗位
	if err := DB.Where("status = ?", "enabled").Find(&positions).Error; err != nil {
		return nil, err
	}

	// 构建岗位映射：部门ID -> 岗位列表
	positionMap := make(map[uint][]models.Position)
	for _, pos := range positions {
		positionMap[pos.DepartmentID] = append(positionMap[pos.DepartmentID], pos)
	}

	// 构建部门的树形结构
	departmentMap := make(map[uint][]models.Department)
	for _, dep := range departments {
		if dep.ParentID != nil {
			// 有父级部门
			departmentMap[*dep.ParentID] = append(departmentMap[*dep.ParentID], dep)
		} else {
			// 没有父级部门，是顶级部门
			departmentMap[dep.CompanyID] = append(departmentMap[dep.CompanyID], dep)
		}
	}

	// 递归构建子部门树，并为每个部门添加岗位
	var buildDepartmentTree func(deps []models.Department, companyID uint) []models.Department
	buildDepartmentTree = func(deps []models.Department, companyID uint) []models.Department {
		var result []models.Department
		for i := range deps {
			if deps[i].CompanyID == companyID { // 确保部门属于当前公司
				// 递归构建子部门
				children := buildDepartmentTree(departmentMap[deps[i].ID], companyID)
				deps[i].Children = children
				deps[i].Positions = positionMap[deps[i].ID] // 将岗位作为叶子节点添加
				result = append(result, deps[i])            // 只将当前公司部门添加到结果中
			}
		}
		return result
	}

	// 将部门分配到对应的公司下
	for i := range companies {
		companies[i].Children = buildDepartmentTree(departmentMap[companies[i].ID], companies[i].ID) // 确保传递公司ID
	}

	return companies, nil
}
