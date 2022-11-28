package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/woodlsy/woodGin/middleware"
	controllers2 "novel/app/controllers"
	middleware2 "novel/app/utils/middleware"
)

func Create() *gin.Engine {
	router := gin.New()

	//router.Use(log.GinLogger())

	router.Use(middleware.GinRecovery(true))
	router.Use(middleware.Cors(map[string]string{"Access-Control-Allow-Headers": "Content-Type,Token,Timestamp,Sign,Platform"}))
	router.Use(middleware2.VerifyTime())
	router.Use(middleware2.GetLogin())

	router.GET("/category", controllers2.Category{}.All)
	router.GET("/category/info", controllers2.Category{}.Info)

	router.GET("/bookIndex", controllers2.Book{}.GetBookListByWeekClick)
	router.GET("/book/list", controllers2.Book{}.List)
	router.GET("/book", controllers2.Book{}.Info)

	router.GET("/articles", controllers2.Article{}.List)
	router.GET("/article", controllers2.Article{}.Info)

	// login
	router.GET("/yzm", controllers2.Login{}.Yzm)
	router.POST("/sms", controllers2.Login{}.SendSmsCode)
	router.POST("/register", controllers2.Login{}.Register)
	router.POST("/login", controllers2.Login{}.Login)

	// member
	router.GET("/member", middleware2.CheckLogin(), controllers2.Member{}.LoginInfo)
	router.GET("/member/book", middleware2.CheckLogin(), controllers2.Member{}.Book)
	router.POST("/member/book/del", middleware2.CheckLogin(), controllers2.Member{}.DelBook)
	router.POST("/member/book/add", middleware2.CheckLogin(), controllers2.Member{}.AddBook)
	router.POST("/article/read", middleware2.CheckLogin(), controllers2.Member{}.Read)

	router.POST("/apply", middleware2.CheckLogin(), controllers2.Member{}.Apply)
	router.GET("/apply/list", controllers2.Member{}.ApplyList)

	router.GET("/config", controllers2.Config{}.Index)

	return router
}
