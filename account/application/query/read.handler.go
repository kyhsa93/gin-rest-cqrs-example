package query

import (
	"errors"

	"github.com/kyhsa93/gin-rest-example/account/domain/model"
)

func (queryBus *QueryBus) handleReadAccountByIDQuery(query *ReadAccountByIDQuery) (*model.Account, error) {
	entity := queryBus.infrastructure.Repository.FindByID(query.AccountID)

	if entity.ID == "" {
		return nil, errors.New("Account is not found")
	}

	return queryBus.entityToModel(entity), nil
}

func (queryBus *QueryBus) handleReadAccountQuery(query *ReadAccountQuery) (*model.Account, error) {
	entity := queryBus.infrastructure.Repository.FindByEmailAndProvider(
		query.Email, query.Provider, query.Unscoped,
	)

	if entity.ID == "" {
		return nil, nil
	}

	if err := compareHashAndPassword(entity.Password, query.Password); err != nil {
		return queryBus.entityToModel(entity), err
	}

	if err := compareHashAndPassword(entity.SocialID, query.SocialID); err != nil {
		return queryBus.entityToModel(entity), err
	}

	return queryBus.entityToModel(entity), nil
}
