package router

import (
	"github.com/gin-gonic/gin"
	"study/controller"
	"study/repository"
)

func Delete(context *gin.Context, repository repository.Repository) {
	controller.Delete(context.Param("id"), repository)
}
