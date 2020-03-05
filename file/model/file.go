package model

import "time"

// File file struct
type File struct {
	ID        string    `json:"id" example:"389df385-ccaa-49c1-aee2-698ba1191857"`
	AccountID string    `json:"accountId" example:"389df385-ccaa-49c1-aee2-698ba1191857"`
	Usage     string    `json:"usage" example:"profile"`
	ImageURL  string    `json:"imageUrl" example:"profile.image_url.com"`
	CreatedAt time.Time `json:"createdAt" example:"2019-12-23 12:27:37"`
}
