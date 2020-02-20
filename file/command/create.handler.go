package command

import "github.com/kyhsa93/gin-rest-cqrs-example/file/model"

func (bus *Bus) handleCreateCommand(command *CreateCommand) (*model.File, error) {
	transaction := bus.repository.TransactionStart()
	s3ObjectKey := bus.aws.S3().Upload(command.Image)
	fileEntity, err := bus.repository.Create(s3ObjectKey, command.AccountID, command.Usage, transaction)
	if err != nil {
		bus.repository.TransactionRollback(transaction)
		bus.aws.S3().Delete(s3ObjectKey)
		panic(err)
	}
	bus.repository.TransactionCommit(transaction)
	return bus.entityToModel(fileEntity), nil
}
