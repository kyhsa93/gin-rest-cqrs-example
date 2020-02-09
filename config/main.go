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
	return &Config{
		Server:   server.New(),
		Database: database.New(),
		Swagger:  swagger.New(),
		Auth:     auth.New(),
		Redis:    redis.New(),
		Email:    email.New(),
		AWS:      aws.New(),
	}
}
