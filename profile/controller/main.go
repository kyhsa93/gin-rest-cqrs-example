package controller

import (
	"github.com/gin-gonic/gin"
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
	config     *config.Config
	api        api.Interface
}

// New create profile controller instance
func New(
	route *gin.Engine,
	commandBus *command.Bus,
	queryBus *query.Bus,
	util *util.Util,
	config *config.Config,
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
}

// AuthenticateHTTPReqeust check http request auth
func (controller *Controller) AuthenticateHTTPReqeust(
	accessToken string, accountID string,
) bool {
	if accessToken == "" || accountID == "" {
		return false
	}
	account, err := controller.api.GetAccountByID(accessToken, accountID)
	if err != nil || account.ID != accountID {
		return false
	}
	return true
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
