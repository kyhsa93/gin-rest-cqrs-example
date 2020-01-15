package config

import (
	"github.com/kyhsa93/go-rest-example/config/auth"
	"github.com/kyhsa93/go-rest-example/config/database"
	"github.com/kyhsa93/go-rest-example/config/server"
	"github.com/kyhsa93/go-rest-example/config/swagger"
)

// Config config stcut
type Config struct {
	swagger  *swagger.Swagger
	Auth     *auth.Auth
	Server   *server.Server
	Database *database.Database
}

// InitializeConfig initialize config
func InitializeConfig() *Config {
	server := server.NewServer()
	database := database.NewDatabase()
	swagger := swagger.NewSwagger()
	auth := auth.New()
	return &Config{Server: server, Database: database, swagger: swagger, Auth: auth}
}
