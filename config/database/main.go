package database

import (
	"github.com/caarlos0/env"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql package for gorm
)

// Database database struct
type Database struct {
	Port     string `env:"DATABASE_PORT" envDefault:"3306"`
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Name     string `env:"DATABASE_NAME" envDefault:"gin-rest-cqrs-example"`
	User     string `env:"DATABASE_USER" envDefault:"root"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"test"`
	Logging  bool   `env:"DATABASE_LOGGING" envDefault:"true"`
}

// New create database instance
func New() *Database {
	database := &Database{}
	env.Parse(database)
	return database
}
