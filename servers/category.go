package servers

import "novel/models"

//
// GetCategoryAll
// @Description: 获取所有分类
// @return []models.Category
//
func GetCategoryAll() []models.Category {
	fields := "id,name"
	return models.Category{}.GetAll(map[string]interface{}{}, "id asc", fields)
}
