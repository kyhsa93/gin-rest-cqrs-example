package command

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/aws"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/email"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/repository"
	"golang.org/x/crypto/bcrypt"
)

// Bus account command
type Bus struct {
	repository repository.Interface
	email      email.Interface
	aws        aws.Interface
}

// New create Bus instance
func New(
	repository repository.Interface,
	email email.Interface,
	aws aws.Interface,
) *Bus {
	return &Bus{repository: repository, email: email, aws: aws}
}

// Handle handle command
func (bus *Bus) Handle(command interface{}) error {
	switch command := command.(type) {
	case *CreateCommand:
		bus.handleCreateCommand(command)
		return nil
	case *UpdateCommand:
		bus.handleUpdateCommand(command)
		return nil
	case *DeleteCommand:
		bus.handleDeleteCommand(command)
		return nil
	default:
		return errors.New("Command is not handled")
	}
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
