package controller

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/api"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

// Controller profile controller
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	queryBus   *query.Bus
	util       *util.Util
	config     config.Interface
	api        api.Interface
}

// New create profile controller instance
func New(
	route *gin.Engine,
	commandBus *command.Bus,
	queryBus *query.Bus,
	util *util.Util,
	config config.Interface,
	api api.Interface,
) *Controller {
	controller := &Controller{
		route:      route,
		commandBus: commandBus,
		queryBus:   queryBus,
		util:       util,
		config:     config,
		api:        api,
	}
	controller.SetupRoutes()
	return controller
}

// SetupRoutes setup profile router
func (controller *Controller) SetupRoutes() {
	controller.route.POST("profiles", func(context *gin.Context) {
		controller.create(context)
	})
	controller.route.GET("profiles/:id", func(context *gin.Context) {
		controller.readByID(context)
	})
	controller.route.GET("profiles", func(context *gin.Context) {
		controller.read(context)
	})
	controller.route.PUT("profiles", func(context *gin.Context) {
		controller.update(context)
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

// ValidateFileID validate image key
func (controller *Controller) ValidateFileID(accountID string, fileID string) bool {
	if fileID == "" {
		return true
	}
	file, err := controller.api.GetFileByID(fileID)
	if err != nil || file.AccountID != accountID {
		return false
	}
	return true
}
