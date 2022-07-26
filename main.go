package main

import (
	"novel/woodlsy/core"
	"novel/woodlsy/driver"
	_ "novel/woodlsy/driver"
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
	driver.OrmInit()

	novel := &http.Server{
		Addr:    "127.0.0.1:8888",
		Handler: routers.Create(),
	}
	fmt.Println("Listen http://127.0.0.1:8888")
	novel.ListenAndServe()
}
