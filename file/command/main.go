package command

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/file/aws"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/repository"
)

// Bus file command bus
type Bus struct {
	repository repository.Interface
	aws        aws.Interface
}

// New create bus instance
func New(repository repository.Interface, aws aws.Interface) *Bus {
	return &Bus{repository: repository, aws: aws}
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
