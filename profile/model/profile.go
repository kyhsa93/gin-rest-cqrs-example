package model

import "time"

// Profile profile model
type Profile struct {
	ID                    string    `json:"id" example:"389df385-ccaa-49c1-aee2-698ba1191857"`
	ImageURL              string    `json:"image_url" example:"profile.image_url.com"`
	Gender                string    `json:"gender" exmaple:"male"`
	InterestedField       string    `json:"interested_field" example:"develop"`
	InterestedFieldDetail []string  `json:"interested_field_detail" example:"web,server"`
	CreatedAt             time.Time `json:"created_at" example:"2019-12-23 12:27:37"`
	UpdatedAt             time.Time `json:"updated_at" example:"2019-12-23 12:27:37"`
}
