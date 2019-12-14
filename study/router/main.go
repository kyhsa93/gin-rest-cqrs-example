package router

import (
	"github.com/kyhsa93/go-rest-example/config"
	"github.com/kyhsa93/go-rest-example/study/repository"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {
	repository := repository.NewRepository(config.GetConnection())

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
}
