package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kacifer/buggy_gin"
	"github.com/surfinggo/mc"
)

func setupRouter() *gin.Engine {
	r := gin.Default()

	buggy_gin.UseAll(r)

	return r
}

func main() {
	r := setupRouter()

	mc.Must(r.Run(":8080"))
}
