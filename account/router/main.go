package router

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-example/account/service"
	"github.com/kyhsa93/gin-rest-example/util"
)

// Router account router struct
type Router struct {
	route   *gin.Engine
	service service.Interface
	util    *util.Util
}

// New create account router instance
func New(route *gin.Engine, service service.Interface, util *util.Util) *Router {
	router := &Router{route: route, service: service, util: util}
	router.SetupRoutes()
	return router
}

// SetupRoutes setup accounts route handler
func (router *Router) SetupRoutes() {
	router.route.POST("accounts", func(context *gin.Context) {
		router.create(context)
	})

	router.route.GET("accounts/:id", func(context *gin.Context) {
		router.readAccountByID(context)
	})

	router.route.GET("accounts", func(context *gin.Context) {
		router.readAccount(context)
	})

	router.route.PUT("accounts/:id", func(context *gin.Context) {
		router.update(context)
	})

	router.route.DELETE("accounts/:id", func(context *gin.Context) {
		router.delete(context)
	})
}

func emailAndProviderValidation(email string, provider string) bool {
	return strings.Contains(email, "@"+provider+".")
}
