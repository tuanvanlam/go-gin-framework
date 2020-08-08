package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-framework/controller"
	"go-gin-framework/service"
)

var (
	videoService    = service.New()
	videoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})

	server.Run(":8080")
}
