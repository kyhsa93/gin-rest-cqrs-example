package router

import (
	"github.com/gin-gonic/gin"
	"study/repositories"
)

func Delete(context *gin.Context, repository repositories.StudyRepository) {
	context.JSON(200, "delete")
}