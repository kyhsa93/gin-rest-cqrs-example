package controller_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-example/account/application"
	"github.com/kyhsa93/gin-rest-example/account/interface/controller"
	"github.com/kyhsa93/gin-rest-example/util"
)

// TestNew test controller's New method
func TestNew(t *testing.T) {
	_, engine := gin.CreateTestContext(httptest.NewRecorder())
	service := &application.Service{}
	util := &util.Util{}
	controllerInstance := controller.New(engine, service, util)
	if controllerInstance == nil {
		t.Error("Can not create controller instance")
	}
}
