package command

import "github.com/kyhsa93/gin-rest-cqrs-example/file/model"

func (bus *Bus) handleDeleteCommand(command *DeleteCommand) (*model.File, error) {
	return nil, nil
}
