package config

import (
	"github.com/kyhsa93/go-rest-example/docs"
)

func init() {
	docs.SwaggerInfo.Title = "Go REST api Example"
	docs.SwaggerInfo.Description = "This is Example for REST api using Go"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:5000"
}
