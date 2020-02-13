package query

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
)

func (bus *Bus) handleReadAccountByIDQuery(query *ReadAccountByIDQuery) (*model.Account, error) {
	entity := bus.repository.FindByID(query.AccountID)

	if entity.ID == "" {
		return nil, errors.New("Account is not found")
	}

	return bus.entityToModel(entity), nil
}

func (bus *Bus) handleReadAccountQuery(query *ReadAccountQuery) (*model.Account, error) {
	entity := bus.repository.FindByEmailAndProvider(
		query.Email, query.Provider, query.Unscoped,
	)

	if entity.ID == "" {
		return nil, nil
	}

	if err := compareHashAndPassword(entity.Password, query.Password); err != nil {
		return bus.entityToModel(entity), err
	}

	if err := compareHashAndPassword(entity.SocialID, query.SocialID); err != nil {
		return bus.entityToModel(entity), err
	}

	return bus.entityToModel(entity), nil
}
