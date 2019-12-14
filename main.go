package main

import (
	"log"

	accountRouter "github.com/kyhsa93/go-rest-example/account/router"
	"github.com/kyhsa93/go-rest-example/config"
	"github.com/kyhsa93/go-rest-example/docs"
	studyRouter "github.com/kyhsa93/go-rest-example/study/router"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	route := gin.Default()
	accountRouter.SetupRoutes(route)
	studyRouter.SetupRoutes(route)

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(route.Run(":" + config.GetService().Port))
}
