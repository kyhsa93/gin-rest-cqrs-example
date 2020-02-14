package command

import (
	"github.com/google/uuid"
)

func (bus *Bus) handleCreateCommand(command *CreateCommand) {
	uuid, _ := uuid.NewRandom()
	imageKey := bus.aws.S3().Upload(command.Image)
	bus.repository.Save(uuid.String(), command.AccountID, command.Usage, imageKey)
}
