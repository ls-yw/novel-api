package servers

import (
	"fmt"
	"novel/models"
)

func GetOne() models.Category {
	a := models.Category{}
	b := a.GetOne(map[string]interface{}{"id": 88}, "id asc")
	fmt.Printf("%+v", b)
	return a
}
