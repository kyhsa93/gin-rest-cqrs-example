package infrastructure

import (
	"github.com/go-redis/redis"
	"github.com/jinzhu/gorm"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure/aws"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure/email"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure/repository"
	"github.com/kyhsa93/gin-rest-example/config"
)

// Infrastructure infra struct
type Infrastructure struct {
	Repository *repository.Repository
	Email      *email.Email
	AWS        *aws.AWS
}

// New create infra instance
func New(databaseConnection *gorm.DB, redisClient *redis.Client, config *config.Config) *Infrastructure {
	return &Infrastructure{
		Repository: repository.New(redisClient, databaseConnection),
		Email:      email.New(config),
		AWS:        aws.New(config),
	}
}
