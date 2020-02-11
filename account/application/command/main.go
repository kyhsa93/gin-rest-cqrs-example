package command

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/infrastructure"
	"golang.org/x/crypto/bcrypt"
)

// CommandBus account command
type CommandBus struct {
	infrastructure *infrastructure.Infrastructure
}

// New create commandBus instance
func New(infrastructure *infrastructure.Infrastructure) *CommandBus {
	return &CommandBus{infrastructure: infrastructure}
}

// Handle handle command
func (commandBus *CommandBus) Handle(command interface{}) error {
	switch command := command.(type) {
	case *CreateCommand:
		commandBus.handleCreateCommand(command)
		return nil
	case *UpdateCommand:
		commandBus.handleUpdateCommand(command)
		return nil
	case *DeleteCommand:
		commandBus.handleDeleteCommand(command)
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
