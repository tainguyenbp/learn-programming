package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Hello(ctx *gin.Context) {
	name := ctx.Param("name")
	ctx.JSON(200, gin.H{"message": fmt.Sprintf("Hello %s!", name)})
}

func Engine() *gin.Engine {
	engine := gin.Default()
	engine.GET("/hello/:name", Hello)
	return engine
}

func main() {
	if err := Engine().Run("0.0.0.0:8088"); err != nil {
		println("ERROR running server:", err.Error())
	}
}
