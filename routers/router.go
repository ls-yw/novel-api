package routers

import (
	"github.com/gin-gonic/gin"
	"novel/controllers"
	middleware2 "novel/utils/middleware"
	"novel/woodlsy/middleware"
)

func Create() *gin.Engine {
	router := gin.New()

	//router.Use(log.GinLogger())

	router.Use(middleware.GinRecovery(true))
	router.Use(middleware.Cors())
	router.Use(middleware2.GetLogin())

	router.GET("/category", controllers.Category{}.All)

	router.GET("/bookIndex", controllers.Book{}.GetBookListByWeekClick)
	router.GET("/book/list", controllers.Book{}.List)
	router.GET("/book", controllers.Book{}.Info)

	router.GET("/articles", controllers.Article{}.List)

	// login
	router.GET("/yzm", controllers.Login{}.Yzm)
	router.POST("/sms", controllers.Login{}.SendSmsCode)
	router.POST("/register", controllers.Login{}.Register)
	router.POST("/login", controllers.Login{}.Login)

	// member
	router.GET("/member", middleware2.CheckLogin(), controllers.Member{}.LoginInfo)
	router.GET("/member/book", middleware2.CheckLogin(), controllers.Member{}.Book)
	router.POST("/member/book/del", middleware2.CheckLogin(), controllers.Member{}.DelBook)
	router.POST("/member/book/add", middleware2.CheckLogin(), controllers.Member{}.AddBook)

	router.POST("/apply", middleware2.CheckLogin(), controllers.Member{}.Apply)
	router.GET("/apply/list", controllers.Member{}.ApplyList)

	return router
}
