package config

import (
	"github.com/kyhsa93/gin-rest-example/config/auth"
	"github.com/kyhsa93/gin-rest-example/config/database"
	"github.com/kyhsa93/gin-rest-example/config/redis"
	"github.com/kyhsa93/gin-rest-example/config/server"
	"github.com/kyhsa93/gin-rest-example/config/swagger"
)

// Config config stcut
type Config struct {
	swagger  *swagger.Swagger
	Auth     *auth.Auth
	Server   *server.Server
	Database *database.Database
	Redis    *redis.Redis
}

// InitializeConfig initialize config
func InitializeConfig() *Config {
	server := server.NewServer()
	database := database.NewDatabase()
	redis := redis.New()
	swagger := swagger.NewSwagger()
	auth := auth.New()
	return &Config{Server: server, Database: database, swagger: swagger, Auth: auth, Redis: redis}
}
