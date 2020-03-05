package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
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
}

// GetAccountByAccessToken check http request auth
func (controller *Controller) GetAccountByAccessToken(
	accessToken string,
) (model.Account, error) {
	if accessToken == "" {
		return model.Account{}, errors.New("token is empty")
	}
	account, err := controller.api.GetAccountByAccessToken(
		accessToken,
	)
	if account == nil || err != nil {
		return model.Account{}, errors.New("token is invalid")
	}
	return *account, nil
}
