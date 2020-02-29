package command

import (
	"github.com/google/uuid"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
)

func (bus *Bus) handleCreateCommand(command *CreateCommand) (*model.Account, error) {
	uuid, _ := uuid.NewRandom()
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(command.Password, command.SocialID)

	createdAccountEntity, createError := bus.repository.Create(
		uuid.String(),
		command.Email,
		command.Provider,
		hashedSocialID,
		hashedPassword,
		command.FileID,
		command.Gender,
		command.InterestedField,
	)
	if createError != nil {
		return nil, createError
	}
	bus.email.Send([]string{command.Email}, "Account is created.")
	accountModel := bus.entityToModel(createdAccountEntity)
	accountModel.AccessToken = accountModel.CreateAccessToken()
	return accountModel, nil
}
