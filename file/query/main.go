package query

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/config"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/entity"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/model"
	"github.com/kyhsa93/gin-rest-cqrs-example/file/repository"
)

// Bus file query bus
type Bus struct {
	config     *config.Config
	repository repository.Interface
}

// New create Bus instance
func New(config *config.Config, repository repository.Interface) *Bus {
	return &Bus{config: config, repository: repository}
}

// Handle handle query
func (bus *Bus) Handle(query interface{}) (*model.File, error) {
	switch query := query.(type) {
	case *ReadFileByIDQuery:
		return bus.handleReadFileByIDQuery(query)
	default:
		return nil, errors.New("Query can not handled")
	}
}

func (bus *Bus) entityToModel(entity entity.File) *model.File {
	var fileModel model.File
	fileModel.ID = entity.ID
	fileModel.AccountID = entity.AccountID
	fileModel.Usage = entity.Usage
	fileModel.CreatedAt = entity.CreatedAt
	fileModel.ImageURL = bus.config.AWS.S3().Endpoint() + "/" +
		bus.config.AWS.S3().Bucket() + "/" + entity.ID
	return &fileModel
}
