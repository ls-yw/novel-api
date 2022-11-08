package models

type Config struct {
	Model
	Type        string `json:"type,omitempty"`
	ConfigKey   string `json:"config_key,omitempty"`
	ConfigValue string `json:"config_value,omitempty"`
}

func (m Config) GetAll(where map[string]interface{}, orderBy string, fields string) []Config {
	list := make([]Config, 0)
	getAll(&list, where, orderBy, fields)
	return list
}
