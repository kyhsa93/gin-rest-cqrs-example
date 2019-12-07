package router

import (
	"github.com/gin-gonic/gin"
	"study/repositories"
)

func Update(context *gin.Context, repository repositories.StudyRepository) {
	context.JSON(200, "update")
}
