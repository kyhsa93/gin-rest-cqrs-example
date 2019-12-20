package router

import (
	"github.com/gin-gonic/gin"
)

// SetupRoutes setup accounts route handler
func SetupRoutes(route *gin.Engine) {
	route.POST("accounts", func(context *gin.Context) {
		Create(context)
	})

	route.GET("accounts/:id", func(context *gin.Context) {
		ReadItem(context)
	})

	route.GET("accounts", func(context *gin.Context) {
		ReadList(context)
	})

	route.PUT("accounts/:id", func(context *gin.Context) {
		Update(context)
	})

	route.DELETE("accounts/:id", func(context *gin.Context) {
		Delete(context)
	})
}
