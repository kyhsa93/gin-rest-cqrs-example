package controller

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/api"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

// Controller account controller struct
type Controller struct {
	route      *gin.Engine
	commandBus *command.Bus
	queryBus   *query.Bus
	util       *util.Util
	config     *config.Config
	api        api.Interface
}

// New create account controller instance
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

// AuthenticateHTTPRequest check http request auth
func (controller *Controller) AuthenticateHTTPRequest(context *gin.Context) {
	accessToken := context.GetHeader("Authorization")
	if accessToken == "" {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	id := context.Param("id")
	if id == "" {
		httpError := controller.util.Error.HTTP.BadRequest()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	query := &query.ReadAccountByIDQuery{AccountID: id}
	account, queryError := controller.queryBus.Handle(query)
	if queryError != nil {
		httpError := controller.util.Error.HTTP.Unauthorized()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}

	account.AccessToken = accessToken
	if !account.ValidateAccessToken() {
		httpError := controller.util.Error.HTTP.Forbidden()
		context.JSON(httpError.Code(), httpError.Message())
		return
	}
	return
}

// ValidateFileID validate image key
func (controller *Controller) ValidateFileID(accountID string, fileID string) bool {
	file, err := controller.api.GetFileByID(fileID)
	if err != nil || file.AccountID != accountID {
		return false
	}
	return true
}

func emailAndProviderValidation(email string, provider string) bool {
	if provider == "email" {
		return true
	}
	return strings.Contains(email, "@"+provider+".")
}
