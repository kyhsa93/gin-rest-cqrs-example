package command

func (commandBus *CommandBus) handleDeleteCommand(command *DeleteCommand) {
	commandBus.repository.Delete(command.AccountID)
}
