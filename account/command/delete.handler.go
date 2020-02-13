package command

func (bus *Bus) handleDeleteCommand(command *DeleteCommand) {
	bus.repository.Delete(command.AccountID)
}
