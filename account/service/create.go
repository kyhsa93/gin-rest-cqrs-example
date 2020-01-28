package service

import (
	"github.com/google/uuid"
)

// Create create account
func (service *Service) Create(email string, provider string, socialID string, password string) {
	uuid, _ := uuid.NewRandom()
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(password, socialID)
	service.repository.Save(uuid.String(), email, provider, hashedSocialID, hashedPassword)
	service.config.Email().Send([]string{email})
}
