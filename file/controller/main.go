package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

// Controller file controller struct
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	util       *util.Util
}

// New create file controller instance
func New(
	route *gin.Engine,
	commandBus *command.Bus,
	util *util.Util,
) *Controller {
	controller := &Controller{route: route, commandBus: commandBus, util: util}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup files route handler
func (controller *Controller) SetupRoutes() {
	controller.route.POST("/files", func(context *gin.Context) {
		controller.create(context)
	})

	controller.route.PUT("/files/:id", func(context *gin.Context) {
		controller.update(context)
	})

	controller.route.DELETE("/files/:id", func(context *gin.Context) {
		controller.delete(context)
	})
}
