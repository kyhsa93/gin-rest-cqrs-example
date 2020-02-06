package swagger

import (
	"github.com/kyhsa93/gin-rest-example/docs"
)

// Swagger swagger config struct
type Swagger struct{}

// New create swagger configuration instance
func New() *Swagger {
	docs.SwaggerInfo.Title = "Gin REST api Example"
	docs.SwaggerInfo.Description = "This is Example for REST api using Gin"
	docs.SwaggerInfo.Version = "1.0.0"
	docs.SwaggerInfo.Host = "localhost:5000"
	return &Swagger{}
}
