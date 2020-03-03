package config

import (
	"github.com/caarlos0/env"
	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql package for gorm
)

// Database database struct
type Database struct {
	Port     string `env:"DATABASE_PORT" envDefault:"27017"`
	Host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	Name     string `env:"DATABASE_NAME" envDefault:"gin-rest-cqrs-example"`
	User     string `env:"DATABASE_USER" envDefault:"root"`
	Password string `env:"DATABASE_PASSWORD" envDefault:"test"`
}

// NewDatabase create database instance
func NewDatabase() *Database {
	database := &Database{}
	env.Parse(database)
	return database
}
