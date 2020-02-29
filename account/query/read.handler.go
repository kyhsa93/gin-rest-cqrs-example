package query

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/account/model"
)

func (bus *Bus) handleReadAccountByIDQuery(query *ReadAccountByIDQuery) (*model.Account, error) {
	entity := bus.repository.FindByID(query.AccountID, false)

	if entity.ID == "" {
		return nil, errors.New("Account is not found")
	}

	model := bus.entityToModel(entity)
	model.AccessToken = model.CreateAccessToken()
	return model, nil
}

func (bus *Bus) handleReadAccountQuery(query *ReadAccountQuery) (*model.Account, error) {
	entity := bus.repository.FindByEmailAndProvider(
		query.Email, query.Provider, query.Unscoped,
	)

	if entity.ID == "" {
		return &model.Account{}, nil
	}

	if err := compareHashAndPassword(entity.Password, query.Password); err != nil {
		return &model.Account{}, err
	}

	if err := compareHashAndPassword(entity.SocialID, query.SocialID); err != nil {
		return &model.Account{}, err
	}

	model := bus.entityToModel(entity)
	model.AccessToken = model.CreateAccessToken()
	return model, nil
}
