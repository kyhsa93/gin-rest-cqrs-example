package command

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
)

func (bus *Bus) handleUpdateCommand(
	command *UpdateCommand,
) (*model.Account, error) {
	oldData := bus.repository.FindByID(command.AccountID, false)
	if oldData.ID == "" {
		return nil, errors.New("Update target Account data is not found")
	}
	hashedPassword, hashedSocialID :=
		getHashedPasswordAndSocialID(command.Password, command.SocialID)

	updatedAccountEntity, updateError := bus.repository.Update(
		oldData.ID,
		command.Email,
		command.Provider,
		hashedSocialID,
		hashedPassword,
	)
	if updateError != nil {
		return nil, updateError
	}

	bus.email.Send([]string{command.Email}, "Account is updated.")
	accountModel := bus.entityToModel(updatedAccountEntity)
	accountModel.CreateAccessToken()
	return accountModel, nil
}
