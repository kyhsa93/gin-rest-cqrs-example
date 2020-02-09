package command

import (
	"github.com/google/uuid"
)

func (commandBus *CommandBus) handleCreateCommand(command *CreateCommand) {
	uuid, _ := uuid.NewRandom()
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(command.Password, command.SocialID)

	imageKey := ""
	if command.Image != nil {
		imageKey = commandBus.infrastructure.AWS.S3.Upload(command.Image)
	}

	commandBus.infrastructure.Repository.Save(
		uuid.String(),
		command.Email,
		command.Provider,
		hashedSocialID,
		hashedPassword,
		imageKey,
		command.Gender,
		command.Intereste,
	)

	commandBus.infrastructure.Email.Send([]string{command.Email}, "Account is created.")
}
