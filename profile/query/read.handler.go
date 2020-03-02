package query

import (
	"errors"

	"github.com/kyhsa93/gin-rest-cqrs-example/profile/model"
)

func (bus *Bus) handleReadProfileByIDQuery(
	query *ReadProfileByIDQuery,
) (*model.Profile, error) {
	profileEntity, err := bus.repository.FindByID(query.ProfileID)
	if err != nil {
		return nil, err
	}

	if profileEntity.ID == "" {
		return nil, errors.New("Profile is not found")
	}
	return bus.entityToModel(profileEntity), nil
}
