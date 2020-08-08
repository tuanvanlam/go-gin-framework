package main

import (
	"github.com/gin-gonic/gin"
	gindump "github.com/tpkeeper/gin-dump"
	"go-gin-framework/controller"
	"go-gin-framework/middlewares"
	"go-gin-framework/service"
	"io"
	"net/http"
	"os"
)

var (
	videoService    = service.New()
	videoController = controller.New(videoService)
)

func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {

	setupLogOutput()

	server := gin.New()

	server.Use(gin.Recovery(), middlewares.Logger(), middlewares.BasicAuth(), gindump.Dump())

	server.GET("/videos", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindAll())
	})

	server.POST("/videos", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Video Input is Valid!!"})
		}
	})

	_ = server.Run(":8080")
}
