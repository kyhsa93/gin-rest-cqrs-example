package router

import (
	"study/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(repository *repositories.Repository) *gin.Engine {
	route := gin.Default()

	route.POST("studies", func(context *gin.Context) {
		Create(context, *repository)
	})

	route.GET("studies/:id", func(context *gin.Context) {
		ReadItem(context, *repository)
	})

	route.GET("studies", func(context *gin.Context) {
		ReadList(context, *repository)
	})

	route.PUT("studies/:id", func(context *gin.Context) {
		Update(context, *repository)
	})

	route.DELETE("studies/:id", func(context *gin.Context) {
		Delete(context, *repository)
	})

	return route
}
