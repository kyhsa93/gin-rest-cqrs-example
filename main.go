package main

import (
	"log"
	"study/config"
	"study/docs"
	"study/model"
	"study/repository"
	"study/router"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	serviceConfig := config.GetService()
	databaseConfig := config.GetDatabase()

	databaseUser := databaseConfig.User
	databasePassword := databaseConfig.Password
	databaseName := databaseConfig.Name
	databaseHost := databaseConfig.Host
	databasePort := databaseConfig.Port

	db, err := gorm.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/"+databaseName+"?parseTime=true")

	if err != nil {
		log.Println(err)
		panic(err)
	}

	db.AutoMigrate(model.Study{})

	db.LogMode(true)
	defer db.Close()

	port := serviceConfig.Port
	repository := repository.NewRepository(db)
	route := router.SetupRoutes(repository)

	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "This is a sample server for Swagger."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:5000"

	route.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	log.Fatal(route.Run(":" + port))
}
