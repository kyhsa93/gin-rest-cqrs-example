package command

import "mime/multipart"

// CreateCommand create file command
type CreateCommand struct {
	AccountID   string
	AccessToken string
	Usage       string
	Image       *multipart.FileHeader
}
