package main

import (
	"novel/woodlsy/core"
	"novel/woodlsy/db"
	"novel/woodlsy/log"
)

import (
	"fmt"
	"net/http"
	"novel/routers"
)

func main() {

	core.ConfigInit()
	log.LogInit()
	core.RedisInit()
	db.OrmInit()

	novel := &http.Server{
		Addr:    ":7001",
		Handler: routers.Create(),
	}
	fmt.Println("Listen http://127.0.0.1:7001")
	novel.ListenAndServe()
}
