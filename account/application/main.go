package application

import (
	"mime/multipart"

	"github.com/kyhsa93/gin-rest-example/account/domain/model"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure"
	"github.com/kyhsa93/gin-rest-example/account/infrastructure/entity"
	"github.com/kyhsa93/gin-rest-example/config"
	"golang.org/x/crypto/bcrypt"
)

// Interface service interface
type Interface interface {
	Create(
		email string,
		provider string,
		socialID string,
		password string,
		image *multipart.FileHeader,
		gender string,
		intereste string,
	)
	ReadAccountByID(acountID string) *model.Account
	ReadAccount(email string,
		provider string,
		socialID string,
		password string,
		unscoped bool,
	) (*model.Account, error)
	Update(
		accountID string,
		email string,
		provider string,
		socialID string,
		password string,
		image *multipart.FileHeader,
		gender string,
		intereste string,
	)
	Delete(accountID string)
}

// Service account service struct
type Service struct {
	infrastructure *infrastructure.Infrastructure
	config         *config.Config
}

func (service *Service) entityToModel(entity entity.Account) *model.Account {
	var accountModel model.Account
	accountModel.ID = entity.ID
	accountModel.Email = entity.Email
	accountModel.Provider = entity.Provider
	accountModel.Gender = entity.Gender
	accountModel.ImageURL = service.config.AWS.S3.Endpoint + "/" + service.config.AWS.S3.Bucket + "/" + entity.ImageKey
	accountModel.Intereste = entity.Intereste
	accountModel.CreatedAt = entity.CreatedAt
	accountModel.UpdatedAt = entity.UpdatedAt
	return &accountModel
}

func getHashedPasswordAndSocialID(password string, socialID string) (string, string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}

	hashedSocialID, err := bcrypt.GenerateFromPassword([]byte(socialID), bcrypt.MinCost)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword), string(hashedSocialID)
}

// New create account service instance
func New(infrastructure *infrastructure.Infrastructure, config *config.Config) *Service {
	return &Service{infrastructure: infrastructure, config: config}
}
