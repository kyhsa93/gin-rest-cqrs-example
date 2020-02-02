package aws

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/caarlos0/env"
	"github.com/google/uuid"
	s3Config "github.com/kyhsa93/gin-rest-example/config/aws/s3"
)

// Interface aws config interfcae
type Interface interface {
	AddFileToS3(fileHeader *multipart.FileHeader) string
}

type awsEnvironmentValue struct {
	SecretID  string `env:"AWS_SECRET_ID" envDefault:"aws_secret_id"`
	SecretKey string `env:"AWS_SECRET_KEY" envDefault:"aws_secret_key"`
	Token     string `env:"AWS_TOKEN" envDefault:"aws_token"`
}

// AWS default aws config
type AWS struct {
	awsEnvironmentValue *awsEnvironmentValue
	S3                  *s3Config.S3
}

// New create aws config instance
func New() *AWS {
	awsENV := &awsEnvironmentValue{}
	config := &AWS{}
	env.Parse(awsENV)
	config.awsEnvironmentValue = awsENV
	config.S3 = s3Config.New()
	return config
}

func (awsConfig *AWS) awsEndpointResolver(
	service,
	region string,
	optFns ...func(*endpoints.Options),
) (endpoints.ResolvedEndpoint, error) {
	if service == endpoints.S3ServiceID {
		return endpoints.ResolvedEndpoint{
			URL:           awsConfig.S3.Endpoint,
			SigningRegion: awsConfig.S3.Region,
		}, nil
	}

	return endpoints.DefaultResolver().EndpointFor(service, region, optFns...)
}

func (awsConfig *AWS) getAWSSession() *session.Session {
	return session.Must(session.NewSession(&aws.Config{
		Region:           aws.String(endpoints.ApNortheast2RegionID),
		EndpointResolver: endpoints.ResolverFunc(awsConfig.awsEndpointResolver),
		Credentials: credentials.NewStaticCredentials(
			awsConfig.awsEnvironmentValue.SecretID,
			awsConfig.awsEnvironmentValue.SecretKey,
			awsConfig.awsEnvironmentValue.Token,
		),
		S3ForcePathStyle: aws.Bool(true),
	}))
}

// AddFileToS3 upload file to aws s3 and return file key
func (awsConfig *AWS) AddFileToS3(fileHeader *multipart.FileHeader) string {
	if fileHeader == nil {
		return ""
	}
	awsSession := awsConfig.getAWSSession()

	var size int64 = fileHeader.Size
	buffer := make([]byte, size)
	file, _ := fileHeader.Open()
	file.Read(buffer)

	s3ObjectKey, s3ObjectKeyError := uuid.NewRandom()
	s3ObjectKeyString := s3ObjectKey.String()
	if s3ObjectKeyError != nil {
		s3ObjectKeyString = ""
		log.Println(s3ObjectKeyError)
	}

	_, s3PutObjectError := s3.New(awsSession).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(awsConfig.S3.Bucket),
		Key:                  aws.String(s3ObjectKeyString),
		ACL:                  aws.String(awsConfig.S3.ACL),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String(awsConfig.S3.ContentDispositoin),
		ServerSideEncryption: aws.String(awsConfig.S3.ServerSideEncryption),
		StorageClass:         aws.String(awsConfig.S3.StorageClass),
	})
	if s3PutObjectError != nil {
		s3ObjectKeyString = ""
		log.Println(s3PutObjectError)
	}
	return s3ObjectKeyString
}
