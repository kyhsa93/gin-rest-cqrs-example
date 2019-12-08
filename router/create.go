package router

import (
	"github.com/gin-gonic/gin"
	"study/controller"
	"study/model"
	"study/repository"
)

func Create(context *gin.Context, repository repository.Repository) {
	var study model.Study
	context.ShouldBindJSON(&study)
	controller.Create(&study, repository)
}
