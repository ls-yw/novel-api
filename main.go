package main

import (
	"net/http"
	"novel/routers"
)

func main() {
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	novel := &http.Server{
		Addr:    "127.0.0.1:8888",
		Handler: routers.Init(),
	}
	novel.ListenAndServe()
}
