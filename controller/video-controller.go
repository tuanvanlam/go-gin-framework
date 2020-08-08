package controller

import (
	"github.com/gin-gonic/gin"
	"go-gin-framework/entity"
	"go-gin-framework/service"
	"go-gin-framework/validators"
	"gopkg.in/go-playground/validator.v9"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() []entity.Video
}

type controller struct {
	service service.VideoService
}

var (
	validate *validator.Validate
)

func New(service service.VideoService) VideoController {
	validate = validator.New()
	_ = validate.RegisterValidation("is-cool", validators.ValidateCoolTitle)
	return &controller{service: service}
}

func (controller *controller) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller *controller) Save(ctx *gin.Context) error {
	var video entity.Video

	err := ctx.ShouldBindJSON(&video)
	if err != nil {
		return err
	}

	err = validate.Struct(video)
	if err != nil {
		return err
	}

	controller.service.Save(video)
	return nil
}
