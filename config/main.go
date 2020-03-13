package config

// Config config stcut
type Config struct {
	swagger  *Swagger
	auth     AuthConfigInterface
	server   ServerConfigInterface
	database DatabaseConfigInterface
	redis    RedisConfiginterface
	email    EmailConfigInterface
	aWS      AWSConfigInterface
}

// Interface config interface
type Interface interface {
	Swagger() *Swagger
	Auth() AuthConfigInterface
	Server() ServerConfigInterface
	Database() DatabaseConfigInterface
	Redis() RedisConfiginterface
	Email() EmailConfigInterface
	AWS() AWSConfigInterface
}

// Initialize initialize config
func Initialize() Interface {
	return &Config{
		server:   NewServerConfig(),
		database: NewDatabaseConfig(),
		swagger:  NewSwagger(),
		auth:     NewAuthConfig(),
		redis:    NewRedisConfig(),
		email:    NewEmailConfig(),
		aWS:      NewAWSConfig(),
	}
}

// Swagger get swagger config
func (config *Config) Swagger() *Swagger {
	return config.swagger
}

// Auth get auth config
func (config *Config) Auth() AuthConfigInterface {
	return config.auth
}

// Server get server config
func (config *Config) Server() ServerConfigInterface {
	return config.server
}

// Database get database config
func (config *Config) Database() DatabaseConfigInterface {
	return config.database
}

// Redis get redis config
func (config *Config) Redis() RedisConfiginterface {
	return config.redis
}

// Email get email config
func (config *Config) Email() EmailConfigInterface {
	return config.email
}

// AWS get aws config
func (config *Config) AWS() AWSConfigInterface {
	return config.aWS
}
