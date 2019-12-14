package router

import (
	"github.com/kyhsa93/go-rest-example/account/repository"
	"github.com/kyhsa93/go-rest-example/config"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(route *gin.Engine) {
	repository := repository.NewRepository(config.GetConnection())

	route.POST("accounts", func(context *gin.Context) {
		Create(context, *repository)
	})

	route.GET("accounts/:id", func(context *gin.Context) {
		ReadItem(context, *repository)
	})

	route.GET("accounts", func(context *gin.Context) {
		ReadList(context, *repository)
	})

	route.PUT("accounts/:id", func(context *gin.Context) {
		Update(context, *repository)
	})

	route.DELETE("accounts/:id", func(context *gin.Context) {
		Delete(context, *repository)
	})
}
