package query

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/entity"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/repository"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"golang.org/x/crypto/bcrypt"
)

// QueryBus account query bus
type QueryBus struct {
	config     *config.Config
	repository repository.Interface
}

// New create queryBus instance
func New(config *config.Config, repository repository.Interface) *QueryBus {
	return &QueryBus{config: config, repository: repository}
}

// Handle handle query
func (queryBus *QueryBus) Handle(query interface{}) (*model.Account, error) {
	switch query := query.(type) {
	case *ReadAccountByIDQuery:
		return queryBus.handleReadAccountByIDQuery(query)
	case *ReadAccountQuery:
		return queryBus.handleReadAccountQuery(query)
	default:
		return nil, errors.New("Query can not handled")
	}
}

func (queryBus *QueryBus) entityToModel(entity entity.Account) *model.Account {
	var accountModel model.Account
	accountModel.ID = entity.ID
	accountModel.Email = entity.Email
	accountModel.Provider = entity.Provider
	accountModel.Gender = entity.Gender
	accountModel.ImageURL = queryBus.config.AWS.S3.Endpoint + "/" + queryBus.config.AWS.S3.Bucket + "/" + entity.ImageKey
	accountModel.Interest = entity.Interest
	accountModel.CreatedAt = entity.CreatedAt
	accountModel.UpdatedAt = entity.UpdatedAt
	return &accountModel
}

func compareHashAndPassword(hashed string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return err
	}
	return nil
}
