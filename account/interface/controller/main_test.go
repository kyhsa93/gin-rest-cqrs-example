package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/application/command"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/application/query"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/interface/controller"
	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

// TestNew test controller's New method
func TestNew(t *testing.T) {
	_, engine := gin.CreateTestContext(httptest.NewRecorder())
	util := &util.Util{}
	commandBus := &command.CommandBus{}
	queryBus := &query.QueryBus{}
	controllerInstance := controller.New(engine, commandBus, queryBus, util)
	if controllerInstance == nil {
		t.Error("Can not create controller instance")
	}
}
