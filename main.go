package main

import (
	"log"

	accountRouter "github.com/kyhsa93/go-rest-example/account/router"
	"github.com/kyhsa93/go-rest-example/config"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	route := gin.Default()
	accountRouter.SetupRoutes(route)

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(route.Run(":" + config.GetService().Port))
}
