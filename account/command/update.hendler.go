package command

import "github.com/kyhsa93/gin-rest-cqrs-example/account/model"

func (bus *Bus) handleUpdateCommand(command *UpdateCommand) *model.Account {
	oldData := bus.repository.FindByID(command.AccountID, false)
	if oldData.ID == "" {
		return nil
	}
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(command.Password, command.SocialID)
	imageKey := ""
	if command.Image != nil {
		imageKey = bus.aws.S3().Upload(command.Image)
	}
	updatedAccountEntity := bus.repository.Save(
		oldData.ID,
		command.Email,
		command.Provider,
		hashedSocialID,
		hashedPassword,
		imageKey,
		command.Gender,
		command.Interest,
	)
	bus.email.Send([]string{command.Email}, "Account is created.")
	return bus.entityToModel(*updatedAccountEntity)
}
