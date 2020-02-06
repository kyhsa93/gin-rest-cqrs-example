package config

import (
	"github.com/kyhsa93/gin-rest-example/config/auth"
	"github.com/kyhsa93/gin-rest-example/config/aws"
	"github.com/kyhsa93/gin-rest-example/config/database"
	"github.com/kyhsa93/gin-rest-example/config/email"
	"github.com/kyhsa93/gin-rest-example/config/redis"
	"github.com/kyhsa93/gin-rest-example/config/server"
	"github.com/kyhsa93/gin-rest-example/config/swagger"
)

// Config config stcut
type Config struct {
	Swagger  *swagger.Swagger
	Auth     *auth.Auth
	Server   *server.Server
	Database *database.Database
	Redis    *redis.Redis
	Email    *email.Email
	AWS      *aws.AWS
}

// InitializeConfig initialize config
func InitializeConfig() *Config {
	server := server.New()
	database := database.New()
	redis := redis.New()
	swagger := swagger.New()
	auth := auth.New()
	email := email.New()
	aws := aws.New()
	return &Config{
		Server:   server,
		Database: database,
		Swagger:  swagger,
		Auth:     auth,
		Redis:    redis,
		Email:    email,
		AWS:      aws,
	}
}
