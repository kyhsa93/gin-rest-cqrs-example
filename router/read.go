package router

import (
	"github.com/gin-gonic/gin"
	"study/controller"
	"study/repository"
)

func ReadItem(context *gin.Context, repository repository.Repository) {
	context.JSON(200, controller.ReadItem(context.Param("id"), repository))
}

func ReadList(context *gin.Context, repository repository.Repository) {
	context.JSON(200, controller.ReadList(repository))
}
