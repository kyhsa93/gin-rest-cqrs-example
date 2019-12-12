package router

import (
	"github.com/gin-gonic/gin"
	"go-rest-example/controller"
	"go-rest-example/repository"
)

func Delete(context *gin.Context, repository repository.Repository) {
	controller.Delete(context.Param("id"), repository)
}
