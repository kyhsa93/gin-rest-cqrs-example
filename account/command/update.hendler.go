package command

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
)

func (bus *Bus) handleUpdateCommand(command *UpdateCommand) (*model.Account, error) {
	oldData := bus.repository.FindByID(command.AccountID, false)
	if oldData.ID == "" {
		return nil, errors.New("Update target Account data is not found")
	}
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(command.Password, command.SocialID)
	imageKey := ""
	if command.Image != nil {
		imageKey = bus.aws.S3().Upload(command.Image)
	}
	transaction := bus.repository.TransactionStart()
	updatedAccountEntity, updateError := bus.repository.Update(
		oldData.ID,
		command.Email,
		command.Provider,
		hashedSocialID,
		hashedPassword,
		imageKey,
		command.Gender,
		command.Interest,
		transaction,
	)
	if updateError != nil {
		bus.repository.TransactionRollback(transaction)
		return nil, updateError
	}
	bus.repository.TransactionCommit(transaction)

	bus.email.Send([]string{command.Email}, "Account is created.")
	return bus.entityToModel(*updatedAccountEntity), nil
}
