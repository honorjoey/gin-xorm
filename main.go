package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/honorjoey/gin-xorm/cache"
	"github.com/honorjoey/gin-xorm/router"
	"log"
)

var (
	port, mode string
)

func init() {
	flag.StringVar(&port, "port", "8080", "server listening on, default 8080")
	flag.StringVar(&mode, "mode", "debug", "server running mode, default debug mode")
}

func main() {
	flag.Parse()

	cache.Init()

	gin.SetMode(mode)
	r := router.Init()
	err := r.Run(":"+port)
	if err != nil {
		log.Fatalf("Start Server %+v\n", err)
	}
}