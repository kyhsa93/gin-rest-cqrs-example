package router

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/go-rest-example/account/service"
)

// Router account router struct
type Router struct {
	route   *gin.Engine
	service *service.Service
}

// NewRouter create account router instance
func NewRouter(route *gin.Engine, service *service.Service) *Router {
	router := &Router{route: route, service: service}
	router.SetupRoutes()
	return router
}

// SetupRoutes setup accounts route handler
func (router *Router) SetupRoutes() {
	router.route.POST("accounts", func(context *gin.Context) {
		router.create(context)
	})

	router.route.GET("accounts/:id", func(context *gin.Context) {
		router.readItem(context)
	})

	router.route.GET("accounts", func(context *gin.Context) {
		router.readList(context)
	})

	router.route.PUT("accounts/:id", func(context *gin.Context) {
		router.update(context)
	})

	router.route.DELETE("accounts/:id", func(context *gin.Context) {
		router.delete(context)
	})
}
