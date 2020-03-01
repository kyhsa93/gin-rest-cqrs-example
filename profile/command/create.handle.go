package command

import (
	"github.com/google/uuid"
	"github.com/kyhsa93/gin-rest-cqrs-example/profile/model"
)

func (bus *Bus) handleCreateCommand(
	command *CreateCommand,
) (*model.Profile, error) {
	uuid, _ := uuid.NewRandom()
	createdProfileEntity, createError := bus.repository.Create(
		uuid.String(),
		command.AccountID,
		command.Email,
		command.Gender,
		command.FileID,
		command.InterestedFeild,
		command.InterestedFieldDetail,
	)
	if createError != nil {
		return nil, createError
	}
	return bus.entityToModel(createdProfileEntity), nil
}
