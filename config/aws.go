package config

import (
	"os"
	"strconv"
)

// AWSConfigInterface aws config interface
type AWSConfigInterface interface {
	SecretID() string
	SecretKey() string
	Token() string
	S3() S3ConfigInterface
}

// AWS default aws config
type AWS struct {
	secretID  string
	secretKey string
	token     string
	s3        S3ConfigInterface
}

// NewAWSConfig create aws config instance
func NewAWSConfig() *AWS {
	secretID := "aws_secret_id"
	secretKey := "aws_secret_key"
	token := "aws_token"

	if env := os.Getenv("AWS_SECRET_ID"); env != "" {
		secretID = env
	}
	if env := os.Getenv("AWS_SECRET_KEY"); env != "" {
		secretKey = env
	}
	if env := os.Getenv("AWS_TOKEN"); env != "" {
		token = env
	}
	config := &AWS{
		secretID:  secretID,
		secretKey: secretKey,
		token:     token,
		s3:        NewS3(),
	}
	return config
}

// SecretID get aws secretid
func (aws *AWS) SecretID() string {
	return aws.secretID
}

// SecretKey get aws secret key
func (aws *AWS) SecretKey() string {
	return aws.secretKey
}

// Token get aws token
func (aws *AWS) Token() string {
	return aws.token
}

// S3 get aws s3 config interface
func (aws *AWS) S3() S3ConfigInterface {
	return aws.s3
}

// S3ConfigInterface aws s3 config interface
type S3ConfigInterface interface {
	Region() string
	Endpoint() string
	Bucket() string
	ACL() string
	ContentDisposition() string
	ServerSideEncryption() string
	StorageClass() string
	S3ForcePathStyle() bool
}

// S3 default aws s3 config
type S3 struct {
	region               string
	endpoint             string
	bucket               string
	acl                  string
	contentDisposition   string
	serverSideEncryption string
	storageClass         string
	s3ForcePathStyle     bool
}

// NewS3 create s3 config instance
func NewS3() *S3 {
	region := "ap-northeast-2"
	endpoint := "http://localhost:4572"
	bucket := "bucket"
	acl := "public-read"
	contentDisposition := "attachment"
	serverSideEncryption := "AES256"
	storageClass := "INTELLIGENT_TIERING"
	s3ForcePathStyle := true

	if env := os.Getenv("AWS_S3_REGION"); env != "" {
		region = env
	}
	if env := os.Getenv("AWS_S3_ENPOINT"); env != "" {
		endpoint = env
	}
	if env := os.Getenv("AWS_S3_BUCKET"); env != "" {
		bucket = env
	}
	if env := os.Getenv("AWS_S3_ACL"); env != "" {
		acl = env
	}
	if env := os.Getenv("AWS_S3_CONTENT_DISPOSITION"); env != "" {
		contentDisposition = env
	}
	if env := os.Getenv("AWS_S3_SERVER_SIDE_ENCRYPTION"); env != "" {
		serverSideEncryption = env
	}
	if env := os.Getenv("AWS_S3_STORAGE_CLASS"); env != "" {
		storageClass = env
	}
	if env := os.Getenv("AWS_S3_FORCE_PATH_STYLE"); env != "" {
		data, err := strconv.ParseBool(env)
		if err != nil {
			panic(err)
		}
		s3ForcePathStyle = data
	}
	config := &S3{
		region:               region,
		endpoint:             endpoint,
		bucket:               bucket,
		acl:                  acl,
		contentDisposition:   contentDisposition,
		serverSideEncryption: serverSideEncryption,
		storageClass:         storageClass,
		s3ForcePathStyle:     s3ForcePathStyle,
	}
	return config
}

// Region get aws s3 region
func (s3 *S3) Region() string {
	return s3.region
}

// Endpoint get aws s3 endpoint
func (s3 *S3) Endpoint() string {
	return s3.endpoint
}

// Bucket get aws s3 bucket
func (s3 *S3) Bucket() string {
	return s3.bucket
}

// ACL get aws s3 acl
func (s3 *S3) ACL() string {
	return s3.acl
}

// ContentDisposition get aws s3 contentDisposition
func (s3 *S3) ContentDisposition() string {
	return s3.contentDisposition
}

// ServerSideEncryption get aws s3 serverSideEncryption
func (s3 *S3) ServerSideEncryption() string {
	return s3.serverSideEncryption
}

// StorageClass get aws s3 storage class
func (s3 *S3) StorageClass() string {
	return s3.storageClass
}

// S3ForcePathStyle get aws s3 s3 force path style
func (s3 *S3) S3ForcePathStyle() bool {
	return s3.s3ForcePathStyle
}
