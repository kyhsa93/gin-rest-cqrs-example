package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/application/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/application/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

// Controller account controller struct
type Controller struct {
	route      *gin.Engine
	commandBus *command.CommandBus
	queryBus   *query.QueryBus
	util       *util.Util
}

// New create account controller instance
func New(
	route *gin.Engine,
	commandBus *command.CommandBus,
	queryBus *query.QueryBus,
	util *util.Util,
) *Controller {
	controller := &Controller{route: route, commandBus: commandBus, queryBus: queryBus, util: util}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup accounts route handler
func (controller *Controller) SetupRoutes() {
	controller.route.POST("accounts", func(context *gin.Context) {
		controller.create(context)
	})

	controller.route.GET("accounts/:id", func(context *gin.Context) {
		controller.readAccountByID(context)
	})

	controller.route.GET("accounts", func(context *gin.Context) {
		controller.readAccount(context)
	})

	controller.route.PUT("accounts/:id", func(context *gin.Context) {
		controller.update(context)
	})

	controller.route.DELETE("accounts/:id", func(context *gin.Context) {
		controller.delete(context)
	})
}

func emailAndProviderValidation(email string, provider string) bool {
	if provider == "email" {
		return true
	}
	return strings.Contains(email, "@"+provider+".")
}
