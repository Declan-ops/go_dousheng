package main

import (
	"github.com/gin-gonic/gin"
	"go_dousheng/router"
)

func main() {

	r := gin.Default()
	router.InitRouter(r)
	r.Run()
}
