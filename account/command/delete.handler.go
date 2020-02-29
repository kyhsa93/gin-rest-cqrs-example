package command

import "github.com/kyhsa93/gin-rest-cqrs-example/account/model"

func (bus *Bus) handleDeleteCommand(command *DeleteCommand) (*model.Account, error) {
	accountEntity := bus.repository.Delete(command.AccountID)
	return bus.entityToModel(accountEntity), nil
}
