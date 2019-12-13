package router

import (
	"github.com/gin-gonic/gin"
	"go-rest-example/study/controller"
	"go-rest-example/study/repository"
)

func Delete(context *gin.Context, repository repository.Repository) {
	controller.Delete(context.Param("id"), repository)
}
