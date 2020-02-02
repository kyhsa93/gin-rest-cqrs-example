package s3

import "github.com/caarlos0/env"

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

// New create s3 config instance
func New() *S3 {
	config := &S3{}
	env.Parse(config)
	return config
}
