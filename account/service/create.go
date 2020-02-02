package service

import (
	"mime/multipart"

	"github.com/google/uuid"
)

// Create create account
func (service *Service) Create(
	email string,
	provider string,
	socialID string,
	password string,
	image *multipart.FileHeader,
	gender string,
	intereste string,
) {
	uuid, _ := uuid.NewRandom()
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(password, socialID)

	imageKey := ""
	if image != nil {
		imageKey = service.config.AWS.AddFileToS3(image)
	}
	service.repository.Save(
		uuid.String(),
		email,
		provider,
		hashedSocialID,
		hashedPassword,
		imageKey,
		gender,
		intereste,
	)
	service.config.Email.Send([]string{email}, "Account is created.")
}
