package model

import "time"

// File file struct
type File struct {
	ID        string    `json:"id" example:"389df385-ccaa-49c1-aee2-698ba1191857"`
	AccountID string    `json:"account_id" example:"389df385-ccaa-49c1-aee2-698ba1191857"`
	Usage     string    `json:"usage" example:"profile"`
	ImageURL  string    `json:"image_url" example:"profile.image_url.com"`
	CreatedAt time.Time `json:"created_at" example:"2019-12-23 12:27:37"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-12-23 12:27:37"`
}
