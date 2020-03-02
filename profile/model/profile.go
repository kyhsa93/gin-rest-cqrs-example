package model

import "time"

// Profile profile model
type Profile struct {
	ID                    string    `json:"id" example:"profileId"`
	AccountID             string    `json:"accountId" example:"accountId"`
	ImageURL              string    `json:"imageUrl" example:"profile.image_url.com"`
	Gender                string    `json:"gender" exmaple:"male"`
	InterestedField       string    `json:"interestedField" example:"develop"`
	InterestedFieldDetail []string  `json:"interestedFieldDetail" example:"web,server"`
	CreatedAt             time.Time `json:"createdAt" example:"2019-12-23 12:27:37"`
	UpdatedAt             time.Time `json:"updatedAt" example:"2019-12-23 12:27:37"`
}
