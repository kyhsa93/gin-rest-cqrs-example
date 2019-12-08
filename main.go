package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"

	"study/config"
	"study/model"
	"study/repositories"
	"study/router"
)

type Study struct {
	gorm.Model
	Name string
}

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
	repository := repositories.NewRepository(db)
	route := router.SetupRoutes(repository)

	log.Fatal(route.Run(":" + port))
}
