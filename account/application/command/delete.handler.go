package command

func (commandBus *CommandBus) handleDeleteCommand(command *DeleteCommand) {
	commandBus.infrastructure.Repository.Delete(command.AccountID)
}
