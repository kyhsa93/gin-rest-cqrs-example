package command

import "mime/multipart"

// CreateCommand create account command
type CreateCommand struct {
	Email    string
	Provider string
	SocialID string
	Password string
	Image    *multipart.FileHeader
	Gender   string
	Interest string
}
