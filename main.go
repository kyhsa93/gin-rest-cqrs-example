package main

import (
	"log"

	"github.com/kyhsa93/gin-rest-example/account"
	"github.com/kyhsa93/gin-rest-example/config"
	"github.com/kyhsa93/gin-rest-example/util"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey AccessToken
// @in header
// @name Authorization

func main() {
	route := gin.Default()
	config := config.InitializeConfig()
	util := util.InitializeUtil()

	account.InitializeAccount(route, config, util)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(route.Run(":" + config.Server.Port))
}
