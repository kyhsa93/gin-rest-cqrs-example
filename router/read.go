package router

import (
	"github.com/gin-gonic/gin"
	"study/repositories"
)

func Read(context *gin.Context, repository repositories.StudyRepository) {
	context.JSON(200, "read item")
}

func ReadList(context *gin.Context, repository repositories.StudyRepository) {
	context.JSON(200, "read list")
}
