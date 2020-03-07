package command

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/file/model"
)

func (bus *Bus) handleCreateCommand(command *CreateCommand) (*model.File, error) {
	s3ObjectKey := bus.aws.S3().Upload(command.Image)
	fileEntity, err := bus.repository.Create(s3ObjectKey, command.AccountID, command.Usage)
	if err != nil {
		bus.aws.S3().Delete(s3ObjectKey)
		panic(err)
	}
	profile, err := bus.api.GetProfileByAccessToken(command.AccessToken)
	if err != nil {
		return nil, err
	}
	if profile.ID == "" {
		return nil, errors.New("profile is not found")
	}
	_, updateError := bus.api.UpdateProfile(
		command.AccessToken,
		fileEntity.ID,
		profile.InterestedField,
		profile.InterestedFieldDetail,
	)
	if updateError != nil {
		return nil, err
	}
	return bus.entityToModel(fileEntity), nil
}
