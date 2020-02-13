package command

import (
	"github.com/google/uuid"
)

func (commandBus *CommandBus) handleCreateCommand(command *CreateCommand) {
	uuid, _ := uuid.NewRandom()
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(command.Password, command.SocialID)

	imageKey := ""
	if command.Image != nil {
		imageKey = commandBus.aws.S3().Upload(command.Image)
	}

	commandBus.repository.Save(
		uuid.String(),
		command.Email,
		command.Provider,
		hashedSocialID,
		hashedPassword,
		imageKey,
		command.Gender,
		command.Interest,
	)

	commandBus.email.Send([]string{command.Email}, "Account is created.")
}
