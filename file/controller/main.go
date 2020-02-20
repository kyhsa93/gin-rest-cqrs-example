package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/api"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

// Controller file controller struct
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	queryBus   *query.Bus
	util       *util.Util
	api        api.Interface
}

// New create file controller instance
func New(
	route *gin.Engine,
	commandBus *command.Bus,
	queryBus *query.Bus,
	util *util.Util,
	api api.Interface,
) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: commandBus,
		queryBus:   queryBus,
		util:       util,
		api:        api,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup files route handler
func (controller *Controller) SetupRoutes() {
	controller.route.POST("/files", func(context *gin.Context) {
		controller.create(context)
	})

	controller.route.GET("/files/:id", func(context *gin.Context) {
		controller.readFileByID(context)
	})

	controller.route.PUT("/files/:id", func(context *gin.Context) {
		controller.update(context)
	})

	controller.route.DELETE("/files/:id", func(context *gin.Context) {
		controller.delete(context)
	})
}

// AuthenticateHTTPReqeust check http request auth
func (controller *Controller) AuthenticateHTTPReqeust(accessToken string, accountID string) bool {
	if accessToken == "" {
		return false
	}
	account, err := controller.api.GetAccountByID(accessToken, accountID)
	if err != nil || account.ID != accountID {
		return false
	}
	return true
}
