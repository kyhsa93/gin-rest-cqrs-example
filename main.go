package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"

	"study/config"
	"study/repositories"
	"study/model"
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

	db, err := gorm.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/"+databaseName)

	if err != nil {
		log.Println(err)
		panic(err)
	}

	db.AutoMigrate(model.Study{})

	db.LogMode(true)
	defer db.Close()

	port := serviceConfig.Port
	studyRepository := repositories.NewStudyRepository(db)
	route := router.SetupRoutes(studyRepository)

	log.Fatal(route.Run(":"+port))
}
