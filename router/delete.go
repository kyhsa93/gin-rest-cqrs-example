package router

import (
	"github.com/gin-gonic/gin"
	"study/repositories"
)

func Delete(context *gin.Context, repository repositories.Repository) {
	context.JSON(200, "delete")
}
