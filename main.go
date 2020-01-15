package main

import (
	"log"

	"github.com/kyhsa93/go-rest-example/account"
	"github.com/kyhsa93/go-rest-example/config"
	"github.com/kyhsa93/go-rest-example/util"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey AccessToken
// @in header
// @name Authorization

// @securityDefinitions.apikey RefreshToken
// @in header
// @name Refresh
func main() {
	route := gin.Default()
	config := config.InitializeConfig()
	util := util.InitializeUtil()

	account.InitializeAccount(route, config, util)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(route.Run(":" + config.Server.Port))
}
