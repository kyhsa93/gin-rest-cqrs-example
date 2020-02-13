package command

import "mime/multipart"

// UpdateCommand update account command
type UpdateCommand struct {
	AccountID string
	Email     string
	Provider  string
	SocialID  string
	Password  string
	Image     *multipart.FileHeader
	Gender    string
	Interest  string
}
