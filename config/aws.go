package config

import (
	"github.com/caarlos0/env"
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
	S3               *S3
}

// NewAWS create aws config instance
func NewAWS() *AWS {
	awsENV := &EnvironmentValue{}
	config := &AWS{}
	env.Parse(awsENV)
	config.EnvironmentValue = awsENV
	config.S3 = NewS3()
	return config
}

// S3 default aws s3 config
type S3 struct {
	Region               string `env:"AWS_S3_REGION" envDefault:"ap-northeast-2"`
	Endpoint             string `env:"AWS_S3_ENPOINT" envDefault:"http://localhost:4572"`
	Bucket               string `env:"AWS_S3_BUCKET" envDefault:"bucket"`
	ACL                  string `env:"AWS_S3_ACL" envDefault:"public-read"`
	ContentDispositoin   string `env:"AWS_S3_CONTENT_DISPOSITION" envDefault:"attachment"`
	ServerSideEncryption string `env:"AWS_S3_SERVER_SIDE_ENCRYPTION" envDefault:"AES256"`
	StorageClass         string `env:"AWS_S3_STORAGE_CLASS" envDefault:"INTELLIGENT_TIERING"`
	S3ForcePathStyle     bool   `env:"AWS_S3_FORCE_PATH_STYLE" envDefault:"true"`
}

// NewS3 create s3 config instance
func NewS3() *S3 {
	config := &S3{}
	env.Parse(config)
	return config
}
