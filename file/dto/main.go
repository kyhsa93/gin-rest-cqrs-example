package dto

import "mime/multipart"

// Usage file uage lable
func Usage() map[string]string {
	return map[string]string{
		"profile": "profile",
	}
}

// File file dto for file command
type File struct {
	AccountID string
	Usage     string
	File      *multipart.FileHeader
}

// ValidateUsage validate file usage label
func (file *File) ValidateUsage() bool {
	if Usage()[file.Usage] == "" {
		return false
	}
	return true
}
