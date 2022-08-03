package servers

import "novel/models"

func GetBookListByWeekClick() []models.Book {
	fields := "id,name,thumb_img,author,intro"
	return models.Book{}.GetList(map[string]interface{}{}, "weekclick desc", 0, 10, fields)
}
