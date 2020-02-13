package command

func (bus *Bus) handleUpdateCommand(command *UpdateCommand) {
	oldData := bus.repository.FindByID(command.AccountID)
	if oldData.ID == "" {
		return
	}
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(command.Password, command.SocialID)
	imageKey := ""
	if command.Image != nil {
		imageKey = bus.aws.S3().Upload(command.Image)
	}
	bus.repository.Save(
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
}
