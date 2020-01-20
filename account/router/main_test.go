package router_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-example/account/router"
	"github.com/kyhsa93/gin-rest-example/account/service"
	"github.com/kyhsa93/gin-rest-example/util"
)

// TestNew test router's New method
func TestNew(t *testing.T) {
	_, engine := gin.CreateTestContext(httptest.NewRecorder())
	service := &service.Service{}
	util := &util.Util{}
	routerInstance := router.New(engine, service, util)
	if routerInstance == nil {
		t.Error("Can not create router instance")
	}
}
