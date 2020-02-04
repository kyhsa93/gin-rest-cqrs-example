package application

import (
	"github.com/kyhsa93/gin-rest-example/account/domain/model"
	"golang.org/x/crypto/bcrypt"
)

func compareHashAndPassword(hashed string, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(password)); err != nil {
		return err
	}
	return nil
}

// ReadAccountByID read account by acountID
func (service *Service) ReadAccountByID(acountID string) *model.Account {
	entity := service.infrastructure.Repository.FindByID(acountID)

	if entity.ID == "" {
		return nil
	}

	return service.entityToModel(entity)
}

// ReadAccount read account list
func (service *Service) ReadAccount(
	email string,
	provider string,
	socialID string,
	password string,
	unscoped bool,
) (*model.Account, error) {
	entity := service.infrastructure.Repository.FindByEmailAndProvider(email, provider, unscoped)

	if entity.ID == "" {
		return nil, nil
	}

	if err := compareHashAndPassword(entity.Password, password); err != nil {
		return service.entityToModel(entity), err
	}

	if err := compareHashAndPassword(entity.SocialID, socialID); err != nil {
		return service.entityToModel(entity), err
	}

	return service.entityToModel(entity), nil
}
