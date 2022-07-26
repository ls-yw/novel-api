package routers

import (
	"github.com/gin-gonic/gin"
	"novel/controllers"
	"novel/woodlsy/log"
	"novel/woodlsy/middleware"
)

func Create() *gin.Engine {
	router := gin.New()

	router.Use(log.GinLogger())

	router.Use(middleware.GinRecovery(true))

	router.GET("/category", controllers.Category{}.All)

	return router
}
