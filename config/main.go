package config

// Config config stcut
type Config struct {
	Swagger  *Swagger
	Auth     AuthConfigInterface
	Server   ServerConfigInterface
	Database DatabaseConfigInterface
	Redis    RedisConfiginterface
	Email    EmailConfigInterface
	AWS      AWSConfigInterface
}

// InitializeConfig initialize config
func InitializeConfig() *Config {
	return &Config{
		Server:   NewServerConfig(),
		Database: NewDatabaseConfig(),
		Swagger:  NewSwagger(),
		Auth:     NewAuthConfig(),
		Redis:    NewRedisConfig(),
		Email:    NewEmailConfig(),
		AWS:      NewAWSConfig(),
	}
}
