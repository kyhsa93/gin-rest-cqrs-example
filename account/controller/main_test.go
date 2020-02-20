package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/api"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/controller"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

// TestNew test controller's New method
func TestNew(t *testing.T) {
	_, engine := gin.CreateTestContext(httptest.NewRecorder())
	util := &util.Util{}
	commandBus := &command.Bus{}
	queryBus := &query.Bus{}
	config := &config.Config{}
	api := &api.API{}
	controllerInstance := controller.New(
		engine, commandBus, queryBus, util, config, api,
	)
	if controllerInstance == nil {
		t.Error("Can not create controller instance")
	}
}
