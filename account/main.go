package account

import (
	"github.com/gin-gonic/gin"
	"github.com/kyhsa93/go-rest-example/account/repository"
	"github.com/kyhsa93/go-rest-example/account/router"
	"github.com/kyhsa93/go-rest-example/account/service"
	"github.com/kyhsa93/go-rest-example/config"
)

// InitializeAccount innitialize account module
func InitializeAccount(engine *gin.Engine, config *config.Config) {
	repository := repository.NewRepository(config)
	service := service.NewService(repository)
	router.NewRouter(engine, service)
}
