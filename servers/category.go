package servers

import (
	"novel/models"
)

func GetOne() models.Category {
	var a models.Category
	a.GetOne()
	return a
	//fmt.Printf("%+v", a)
}
