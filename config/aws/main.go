package aws

import (
	"github.com/caarlos0/env"
	s3Config "github.com/kyhsa93/gin-rest-example/config/aws/s3"
)

// EnvironmentValue aws environment values
type EnvironmentValue struct {
	SecretID  string `env:"AWS_SECRET_ID" envDefault:"aws_secret_id"`
	SecretKey string `env:"AWS_SECRET_KEY" envDefault:"aws_secret_key"`
	Token     string `env:"AWS_TOKEN" envDefault:"aws_token"`
}

// AWS default aws config
type AWS struct {
	EnvironmentValue *EnvironmentValue
	S3               *s3Config.S3
}

// New create aws config instance
func New() *AWS {
	awsENV := &EnvironmentValue{}
	config := &AWS{}
	env.Parse(awsENV)
	config.EnvironmentValue = awsENV
	config.S3 = s3Config.New()
	return config
}
