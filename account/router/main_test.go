package router_test

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-example/account/router"
	"github.com/kyhsa93/gin-rest-example/account/service"
	"github.com/kyhsa93/gin-rest-example/util"
)

// TestNewRouter test router's NewRouter method
func TestNewRouter(t *testing.T) {
	_, engine := gin.CreateTestContext(httptest.NewRecorder())
	service := &service.Service{}
	util := &util.Util{}
	routerInstance := router.NewRouter(engine, service, util)
	if routerInstance == nil {
		t.Error("Can not create router instance")
	}
}
