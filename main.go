package main

import _ "novel/woodlsy/core"

import (
	"net/http"
	"novel/routers"
)

func main() {

	novel := &http.Server{
		Addr:    "127.0.0.1:8888",
		Handler: routers.Create(),
	}
	novel.ListenAndServe()
}
