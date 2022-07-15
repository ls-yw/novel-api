package routers

import "novel/controllers"

type Api struct {
	Category controllers.Category
}

var ApiGroupApp = new(Api)
