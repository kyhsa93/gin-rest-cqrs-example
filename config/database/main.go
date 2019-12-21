package database

import (
	"log"

	account "github.com/kyhsa93/go-rest-example/account/entity"

	"github.com/caarlos0/env"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql package for gorm
)

type databaseEnvironmentValue struct {
	Port     string `env:"DATABASE_PORT" envDefault:"3306"`
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Name     string `env:"DATABASE_NAME" envDefault:"go-rest-example"`
	User     string `env:"DATABASE_USER" envDefault:"root"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"test"`
	Logging  bool   `env:"DATABASE_LOGGING" envDefault:"true"`
}

// Database database struct
type Database struct {
	databaseEnvironmentValue *databaseEnvironmentValue
	Connection               *gorm.DB
}

// NewDatabase create database instance
func NewDatabase() *Database {
	databaseENV := &databaseEnvironmentValue{}
	database := &Database{}
	env.Parse(databaseENV)
	database.getConnection(databaseENV)
	return database
}

func (database *Database) getConnection(databaseENV *databaseEnvironmentValue) {

	User := databaseENV.User
	Password := databaseENV.Password
	Name := databaseENV.Name
	Host := databaseENV.Host
	Port := databaseENV.Port

	db, err := gorm.Open("mysql", User+":"+Password+"@tcp("+Host+":"+Port+")/"+Name+"?parseTime=true")

	if err != nil {
		log.Println(err)
		panic(err)
	}

	db.LogMode(databaseENV.Logging)
	db.AutoMigrate(account.Account{})

	database.Connection = db
}
