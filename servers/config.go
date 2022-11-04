package servers

import "novel/models"

func GetConfigs() []models.Config {
	return models.Config{}.GetAll(map[string]interface{}{}, "id asc", "type,config_key,config_value")
}
