package account

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/gin-rest-example/account/application"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure"
	"github.com/kyhsa93/gin-rest-example/account/interface/controller"
	"github.com/kyhsa93/gin-rest-example/config"
	"github.com/kyhsa93/gin-rest-example/util"
)

// InitializeAccount innitialize account module
func InitializeAccount(engine *gin.Engine, config *config.Config, util *util.Util) {
	// repository := repository.New(config)
	infra := infrastructure.New(config)
	application := application.New(infra, config)
	controller.New(engine, application, util)
}
