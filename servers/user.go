package servers

import (
	"gorm.io/gorm"
	"novel/models"
	"novel/utils/common"
)

func GetUserByMobile(mobile string, fields string) models.User {
	return models.User{}.GetOne(map[string]interface{}{"mobile": mobile}, "id desc", fields)
}

func GetUserById(id uint, fields string) models.User {
	return models.User{}.GetOne(map[string]interface{}{"id": id}, "id desc", fields)
}

func GetUserByUsername(username string, fields string) models.User {
	return models.User{}.GetOne(map[string]interface{}{"username": username}, "id desc", fields)
}

func UpdateUserByLogin(uid uint, ip string) {
	data := map[string]interface{}{
		"last_ip":   ip,
		"last_time": common.Now(),
		"count":     gorm.Expr("count + ?", 1),
	}
	models.User{}.Update(data, map[string]interface{}{"id": uid})
}
