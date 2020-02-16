package command

import "github.com/kyhsa93/gin-rest-cqrs-example/account/model"

func (bus *Bus) handleDeleteCommand(command *DeleteCommand) (*model.Account, error) {
	transaction := bus.repository.TransactionStart()
	accountEntity := bus.repository.Delete(command.AccountID, transaction)
	bus.repository.TransactionCommit(transaction)
	return bus.entityToModel(accountEntity), nil
}
