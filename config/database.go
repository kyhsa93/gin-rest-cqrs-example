package config

import (
	"os"

	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql package for gorm
)

// DatabaseConfigInterface database config interface
type DatabaseConfigInterface interface {
	Host() string
	Port() string
	Name() string
	User() string
	Password() string
}

// Database database config struct
type Database struct {
	port     string `env:"DATABASE_PORT" envDefault:"27017"`
	host     string `env:"DATABASE_HOST" envDefault:"localhost"`
	name     string `env:"DATABASE_NAME" envDefault:"gin-rest-cqrs-example"`
	user     string `env:"DATABASE_USER" envDefault:"root"`
	password string `env:"DATABASE_PASSWORD" envDefault:"test"`
}

// NewDatabaseConfig create database instance
func NewDatabaseConfig() *Database {
	port := "27017"
	host := "localhost"
	name := "gin-rest-cqrs-example"
	user := "root"
	password := "test"

	if env := os.Getenv("DATABASE_PORT"); env != "" {
		port = env
	}
	if env := os.Getenv("DATABASE_HOST"); env != "" {
		host = env
	}
	if env := os.Getenv("DATABASE_NAME"); env != "" {
		name = env
	}
	if env := os.Getenv("DATABASE_USER"); env != "" {
		user = env
	}
	if env := os.Getenv("DATABASE_PASSWORD"); env != "" {
		password = env
	}
	database := &Database{
		port:     port,
		host:     host,
		name:     name,
		user:     user,
		password: password,
	}
	return database
}

// Host get database host
func (database *Database) Host() string {
	return database.host
}

// Port get database port number
func (database *Database) Port() string {
	return database.port
}

// Name get database name
func (database *Database) Name() string {
	return database.name
}

// User get databsae user name
func (database *Database) User() string {
	return database.user
}

// Password get database user password
func (database *Database) Password() string {
	return database.password
}
