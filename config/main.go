package config

import (
	"github.com/kyhsa93/go-rest-example/config/database"
	"github.com/kyhsa93/go-rest-example/config/server"
	"github.com/kyhsa93/go-rest-example/config/swagger"
)

// Config config stcut
type Config struct {
	Server   *server.Server
	Database *database.Database
	swagger  *swagger.Swagger
}

// InitializeConfig initialize config
func InitializeConfig() *Config {
	server := server.NewServer()
	database := database.NewDatabase()
	swagger := swagger.NewSwagger()
	return &Config{Server: server, Database: database, swagger: swagger}
}
