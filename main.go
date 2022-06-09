package main

import (
	"github.com/gin-gonic/gin"
	"go_dousheng/router"
)

func main() {

	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	//mapper.InitMap()
	router.InitRouter(r)
	r.Run()

}
