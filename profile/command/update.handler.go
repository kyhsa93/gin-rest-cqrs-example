package command

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/profile/model"
)

func (bus *Bus) handleUpdateCommand(
	command *UpdateProfileCommand,
) (*model.Profile, error) {
	olddData, err := bus.repository.FindByID(command.ID)
	if olddData.ID == "" || err != nil {
		return nil, errors.New("update target profile data is not found")
	}

	updatedProfileEntity, err := bus.repository.Update(
		command.ID,
		command.InterestedField,
		command.InterestedFieldDetail,
		command.FileID,
	)
	if err != nil {
		return nil, err
	}
	profileModel := bus.entityToModel(updatedProfileEntity)
	return profileModel, nil
}
