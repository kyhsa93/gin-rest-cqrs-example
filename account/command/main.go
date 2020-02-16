package command

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/aws"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/email"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/entity"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/repository"
	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"golang.org/x/crypto/bcrypt"
)

// Bus account command
type Bus struct {
	repository repository.Interface
	email      email.Interface
	aws        aws.Interface
	config     *config.Config
}

// New create Bus instance
func New(
	repository repository.Interface,
	email email.Interface,
	aws aws.Interface,
	config *config.Config,
) *Bus {
	return &Bus{repository: repository, email: email, aws: aws}
}

// Handle handle command
func (bus *Bus) Handle(command interface{}) (*model.Account, error) {
	switch command := command.(type) {
	case *CreateCommand:
		return bus.handleCreateCommand(command)
	case *UpdateCommand:
		return bus.handleUpdateCommand(command)
	case *DeleteCommand:
		return bus.handleDeleteCommand(command)
	default:
		return nil, errors.New("Command is not handled")
	}
}

func (bus *Bus) entityToModel(entity entity.Account) *model.Account {
	var accountModel model.Account
	accountModel.ID = entity.ID
	accountModel.Email = entity.Email
	accountModel.Provider = entity.Provider
	accountModel.Gender = entity.Gender
	accountModel.Interest = entity.Interest
	accountModel.CreatedAt = entity.CreatedAt
	accountModel.UpdatedAt = entity.UpdatedAt

	if entity.ImageKey != "" {
		accountModel.ImageURL = bus.config.AWS.S3.Endpoint + "/" + bus.config.AWS.S3.Bucket + "/" + entity.ImageKey
	}

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
