package config

type Configs struct {
	Databases Databases `json:"databases"`
	Log       Log       `json:"log"`
	Redis     Redis     `json:"redis"`
	App       App       `json:"app"`
	Aliyun    Aliyun    `json:"aliyun"`
	Api       Api       `json:"api"`
}
