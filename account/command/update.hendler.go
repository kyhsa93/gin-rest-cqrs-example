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
	hashedPassword, _ :=
		getHashedPasswordAndSocialID(command.Password, "")

	if command.Password == "" {
		hashedPassword = oldData.Password
	}

	updatedAccountEntity, updateError := bus.repository.Update(
		oldData.ID,
		hashedPassword,
		command.FCMToken,
	)
	if updateError != nil {
		return nil, updateError
	}

	bus.email.Send([]string{oldData.Email}, "Account is updated.")
	accountModel := bus.entityToModel(updatedAccountEntity)
	accountModel.CreateAccessToken()
	return accountModel, nil
}
