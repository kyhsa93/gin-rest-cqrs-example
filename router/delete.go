package router

import (
	"github.com/gin-gonic/gin"
	"study/controller"
	"study/repositories"
)

func Delete(context *gin.Context, repository repositories.Repository) {
	controller.Delete(context.Param("id"), repository)
}
