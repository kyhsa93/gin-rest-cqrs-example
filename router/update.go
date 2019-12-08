package router

import (
	"github.com/gin-gonic/gin"
	"study/repositories"
)

func Update(context *gin.Context, repository repositories.Repository) {
	context.JSON(200, "update")
}
