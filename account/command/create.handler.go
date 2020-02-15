package command

import (
	"github.com/google/uuid"
	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
)

func (bus *Bus) handleCreateCommand(command *CreateCommand) *model.Account {
	uuid, _ := uuid.NewRandom()
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(command.Password, command.SocialID)

	imageKey := ""
	if command.Image != nil {
		imageKey = bus.aws.S3().Upload(command.Image)
	}

	createdAccountEntity := bus.repository.Save(
		uuid.String(),
		command.Email,
		command.Provider,
		hashedSocialID,
		hashedPassword,
		imageKey,
		command.Gender,
		command.Interest,
	)

	bus.email.Send([]string{command.Email}, "Account is created.")
	return bus.entityToModel(*createdAccountEntity)
}
