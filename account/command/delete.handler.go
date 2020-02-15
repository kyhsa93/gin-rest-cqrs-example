package command

import "github.com/kyhsa93/gin-rest-cqrs-example/account/model"

func (bus *Bus) handleDeleteCommand(command *DeleteCommand) *model.Account {
	bus.repository.Delete(command.AccountID)
	accountEntity := bus.repository.FindByID(command.AccountID, true)
	return bus.entityToModel(accountEntity)
}
