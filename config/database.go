package config

import (
	"log"

	account "github.com/kyhsa93/go-rest-example/account/model"
	study "github.com/kyhsa93/go-rest-example/study/model"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Database struct {
	Port     string `env:"DATABASE_PORT" envDefault:"3306"`
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Name     string `env:"DATABASE_NAME" envDefault:"go-rest-example"`
	User     string `env:"DATABASE_USER" envDefault:"root"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"test"`
}

func GetConnection() *gorm.DB {
	database := Database{}
	env.Parse(&database)

	User := database.User
	Password := database.Password
	Name := database.Name
	Host := database.Host
	Port := database.Port

	db, err := gorm.Open("mysql", User+":"+Password+"@tcp("+Host+":"+Port+")/"+Name+"?parseTime=true")

	if err != nil {
		log.Println(err)
		panic(err)
	}

	db.AutoMigrate(study.Study{}, account.Account{})

	db.LogMode(true)
	return db
}
