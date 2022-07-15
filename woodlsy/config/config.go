package config

type Configs struct {
	Databases Databases `json:"databases"`
	Log       Log       `json:"log"`
	Redis     Redis     `json:"redis"`
}
