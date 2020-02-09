package command

func (commandBus *CommandBus) handleUpdateCommand(command *UpdateCommand) {
	oldData := commandBus.infrastructure.Repository.FindByID(command.AccountID)
	if oldData.ID == "" {
		return
	}
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(command.Password, command.SocialID)
	imageKey := ""
	if command.Image != nil {
		imageKey = commandBus.infrastructure.AWS.S3.Upload(command.Image)
	}
	commandBus.infrastructure.Repository.Save(
		oldData.ID,
		command.Email,
		command.Provider,
		hashedSocialID,
		hashedPassword,
		imageKey,
		command.Gender,
		command.Interest,
	)
	commandBus.infrastructure.Email.Send([]string{command.Email}, "Account is created.")
}
