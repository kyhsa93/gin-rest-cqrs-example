package config

import (
	"github.com/caarlos0/env"
	"github.com/kyhsa93/go-rest-example/docs"
)

type Swagger struct {
	Title       string `env:"SWAGGER_TITLE" envDefault:"Go REST api Example"`
	Description string `env:"SWAGGER_DESCRIPTION" envDefault:"This is Example for REST api using Go"`
	Version     string `envDefault:"1.0.0"`
	Host        string `env:"SWAGGER_HOST" envDefault:"localhost:5000"`
}

func init() {
	swagger := Swagger{}
	env.Parse(&swagger)
	docs.SwaggerInfo.Title = swagger.Title
	docs.SwaggerInfo.Description = swagger.Description
	docs.SwaggerInfo.Version = swagger.Version
	docs.SwaggerInfo.Host = swagger.Host
}
