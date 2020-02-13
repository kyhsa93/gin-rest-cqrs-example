package config

// Config config stcut
type Config struct {
	Swagger  *Swagger
	Auth     *Auth
	Server   *Server
	Database *Database
	Redis    *Redis
	Email    *Email
	AWS      *AWS
}

// InitializeConfig initialize config
func InitializeConfig() *Config {
	return &Config{
		Server:   NewServer(),
		Database: NewDatabase(),
		Swagger:  NewSwagger(),
		Auth:     NewAuth(),
		Redis:    NewRedis(),
		Email:    NewEmail(),
		AWS:      NewAWS(),
	}
}
