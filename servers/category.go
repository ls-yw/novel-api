package servers

import (
	"fmt"
	"novel/models"
)

func GetOne() {
	var a models.Category
	a.GetOne()
	fmt.Printf("%+v", a)
}
