package routers

import (
	"github.com/gin-gonic/gin"
	"novel/controllers"
	"novel/woodlsy/middleware"
)

func Create() *gin.Engine {
	router := gin.New()

	//router.Use(log.GinLogger())

	router.Use(middleware.GinRecovery(true))
	router.Use(middleware.Cors())

	router.GET("/category", controllers.Category{}.All)

	router.GET("/bookIndex", controllers.Book{}.GetBookListByWeekClick)
	router.GET("/book/list", controllers.Book{}.List)
	router.GET("/book", controllers.Book{}.Info)

	router.GET("/articles", controllers.Article{}.List)

	return router
}
