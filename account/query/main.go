package query

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/entity"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/repository"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"golang.org/x/crypto/bcrypt"
)

// Bus account query bus
type Bus struct {
	config     config.Interface
	repository repository.Interface
}

// New create Bus instance
func New(config config.Interface, repository repository.Interface) *Bus {
	return &Bus{config: config, repository: repository}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) (*model.Account, error) {
	switch query := query.(type) {
	case *ReadAccountByIDQuery:
		return bus.handleReadAccountByIDQuery(query)
	case *ReadAccountQuery:
		return bus.handleReadAccountQuery(query)
	case *ReadAccountByEmailQuery:
		return bus.handleReadAccountByEmailquery(query)
	default:
		return nil, errors.New("Query can not handled")
	}
}

func (bus *Bus) entityToModel(entity entity.Account) *model.Account {
	var accountModel model.Account
	accountModel.ID = entity.ID
	accountModel.Email = entity.Email
	accountModel.Provider = entity.Provider
	accountModel.CreatedAt = entity.CreatedAt
	accountModel.UpdatedAt = entity.UpdatedAt

	return &accountModel
}

func compareHashAndPassword(hashed string, password string) error {
	if err := bcrypt.CompareHashAndPassword(
		[]byte(hashed), []byte(password),
	); err != nil {
		return err
	}
	return nil
}
