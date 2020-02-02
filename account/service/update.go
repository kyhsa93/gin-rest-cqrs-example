package service

import (
	"mime/multipart"
)

// Update update account by accountID
func (service *Service) Update(
	accountID string,
	email string,
	provider string,
	socialID string,
	password string,
	image *multipart.FileHeader,
	gender string,
) {
	oldData := service.ReadAccountByID(accountID)
	if oldData == nil {
		return
	}
	hashedPassword, hashedSocialID := getHashedPasswordAndSocialID(password, socialID)
	imageKey := ""
	if image != nil {
		imageKey = service.config.AWS.AddFileToS3(image)
	}
	service.repository.Save(accountID, email, provider, hashedSocialID, hashedPassword, imageKey, gender)
	service.config.Email.Send([]string{email}, "Account is updated.")
}
