package config

import (
	"github.com/kyhsa93/gin-rest-example/config/auth"
	"github.com/kyhsa93/gin-rest-example/config/database"
	"github.com/kyhsa93/gin-rest-example/config/email"
	"github.com/kyhsa93/gin-rest-example/config/redis"
	"github.com/kyhsa93/gin-rest-example/config/server"
	"github.com/kyhsa93/gin-rest-example/config/swagger"
)

// Interface config interface
type Interface interface {
	Auth() *auth.Auth
	Server() *server.Server
	Database() *database.Database
	Redis() *redis.Redis
	Email() *email.Email
}

// Config config stcut
type Config struct {
	swagger  *swagger.Swagger
	auth     *auth.Auth
	server   *server.Server
	database *database.Database
	redis    *redis.Redis
	email    *email.Email
}

// Auth config
func (config *Config) Auth() *auth.Auth {
	return config.auth
}

// Server config
func (config *Config) Server() *server.Server {
	return config.server
}

// Database config
func (config *Config) Database() *database.Database {
	return config.database
}

// Redis config
func (config *Config) Redis() *redis.Redis {
	return config.redis
}

// Email config
func (config *Config) Email() *email.Email {
	return config.email
}

// InitializeConfig initialize config
func InitializeConfig() *Config {
	server := server.NewServer()
	database := database.NewDatabase()
	redis := redis.New()
	swagger := swagger.NewSwagger()
	auth := auth.New()
	email := email.New()
	return &Config{
		server:   server,
		database: database,
		swagger:  swagger,
		auth:     auth,
		redis:    redis,
		email:    email,
	}
}
