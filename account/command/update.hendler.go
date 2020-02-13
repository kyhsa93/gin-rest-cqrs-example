package command

func (commandBus *CommandBus) handleUpdateCommand(command *UpdateCommand) {
	oldData := commandBus.repository.FindByID(command.AccountID)
	if oldData.ID == "" {
		return
	}
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(command.Password, command.SocialID)
	imageKey := ""
	if command.Image != nil {
		imageKey = commandBus.aws.S3().Upload(command.Image)
	}
	commandBus.repository.Save(
		oldData.ID,
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
