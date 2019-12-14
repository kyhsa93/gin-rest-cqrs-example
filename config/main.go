package config

type Config struct {
	service *Service
}

func Init() Config {
	service := GetService()
	config := Config{service: service}
	return config
}
