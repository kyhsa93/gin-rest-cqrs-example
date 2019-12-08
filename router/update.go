package router

import (
	"github.com/gin-gonic/gin"
	"study/controller"
	"study/model"
	"study/repository"
)

func Update(context *gin.Context, repository repository.Repository) {
	var study model.Study
	context.ShouldBindJSON(&study)
	controller.Update(context.Param("id"), &study, repository)
}
