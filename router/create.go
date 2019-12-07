package router

import (
	"github.com/gin-gonic/gin"
	"study/controller"
	"study/model"
	"study/repositories"
)

func Create(context *gin.Context, repository repositories.StudyRepository) {
	var study model.Study
	context.ShouldBindJSON(&study)
	controller.Create(&study, repository)
}
