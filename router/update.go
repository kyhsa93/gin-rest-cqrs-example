package router

import (
	"github.com/gin-gonic/gin"
	"study/controller"
	"study/model"
	"study/repositories"
)

func Update(context *gin.Context, repository repositories.Repository) {
	var study model.Study
	context.ShouldBindJSON(&study)
	controller.Update(context.Param("id"), &study, repository)
}
