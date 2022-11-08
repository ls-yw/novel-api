package servers

import "novel/app/models"

//
// GetCategoryAll
// @Description: 获取所有分类
// @return []models.Category
//
func GetCategoryAll() []models.Category {
	fields := "id,name"
	return models.Category{}.GetAll(map[string]interface{}{}, "id asc", fields)
}

func GetCategoryInfo(id int) models.Category {
	return models.Category{}.GetOne(map[string]interface{}{"id": id}, "id asc", "name,seo_name,keyword,description")
}
