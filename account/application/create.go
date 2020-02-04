package application

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
		imageKey = service.infrastructure.AWS.S3.Upload(image)
	}
	service.infrastructure.Repository.Save(
		uuid.String(),
		email,
		provider,
		hashedSocialID,
		hashedPassword,
		imageKey,
		gender,
		intereste,
	)
	service.infrastructure.Email.Send([]string{email}, "Account is created.")
}
