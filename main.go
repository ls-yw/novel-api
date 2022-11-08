package main

import (
	"fmt"
	"github.com/woodlsy/woodGin/config"
	"github.com/woodlsy/woodGin/db"
	"github.com/woodlsy/woodGin/log"
	"github.com/woodlsy/woodGin/redis"
	"net/http"
	"novel/app/routers"
)

func main() {

	config.Enabled()
	log.Enabled()
	redis.Enabled()
	db.OrmInit()

	novel := &http.Server{
		Addr:    ":7001",
		Handler: routers.Create(),
	}
	fmt.Println("Listen http://127.0.0.1:7001")
	novel.ListenAndServe()
}
