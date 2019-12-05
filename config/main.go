package config

type Config struct {
	service  *Service
	database *Database
}

func Init() Config {
	service := GetService()
	database := GetDatabase()
	config := Config{service: service, database: database}
	return config
}
