package main

import (
	"log"

	"github.com/kyhsa93/go-rest-example/account"
	"github.com/kyhsa93/go-rest-example/config"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	route := gin.Default()
	config := config.InitializeConfig()

	account.InitializeAccount(route, config)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(route.Run(":" + config.Server.Port))
}
