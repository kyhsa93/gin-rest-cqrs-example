package command

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/aws"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/entity"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/model"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/repository"
)

// Bus file command bus
type Bus struct {
	repository repository.Interface
	aws        aws.Interface
	config     *config.Config
}

// New create bus instance
func New(
	repository repository.Interface,
	aws aws.Interface,
	config *config.Config,
) *Bus {
	return &Bus{repository: repository, aws: aws, config: config}
}

// Handle handle command
func (bus *Bus) Handle(command interface{}) (*model.File, error) {
	switch command := command.(type) {
	case *CreateCommand:
		return bus.handleCreateCommand(command)
	default:
		return nil, errors.New("Command is not handled")
	}
}

func (bus *Bus) entityToModel(entity entity.File) *model.File {
	var fileModel model.File
	fileModel.ID = entity.ID
	fileModel.AccountID = entity.AccountID
	fileModel.Usage = entity.Usage
	fileModel.CreatedAt = entity.CreatedAt
	fileModel.UpdatedAt = entity.UpdatedAt
	imageURL := bus.config.AWS.S3.Endpoint + "/" + bus.config.AWS.S3.Bucket + "/" + entity.ID
	fileModel.ImageURL = imageURL

	return &fileModel
}
