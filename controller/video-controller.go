package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-framework/entity"
	"go-gin-framework/service"
)

type VideoController interface {
	Save(ctx *gin.Context) entity.Video
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{service: service}
}

func (controller *controller) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *controller) Save(ctx *gin.Context) entity.Video {
	var video entity.Video

	_ = ctx.BindJSON(&video)
	controller.service.Save(video)
	return video
}
