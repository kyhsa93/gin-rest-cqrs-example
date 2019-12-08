package router

import (
	"github.com/gin-gonic/gin"
	"study/controller"
	"study/repositories"
)

func ReadItem(context *gin.Context, repository repositories.Repository) {
	context.JSON(200, controller.ReadItem(context.Param("id"), repository))
}

func ReadList(context *gin.Context, repository repositories.Repository) {
	context.JSON(200, controller.ReadList(repository))
}
