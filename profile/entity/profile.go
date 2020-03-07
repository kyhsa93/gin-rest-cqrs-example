package entity

import "time"

// Profile profile entity
type Profile struct {
	ID                    string     `json:"_id" bson:"_id"`
	AccountID             string     `json:"accountId" bson:"accountId"`
	Email                 string     `json:"email" bson:"email"`
	Gender                string     `json:"gender" bson:"gender"`
	FileID                string     `json:"fileId" bson:"fileId"`
	InterestedField       string     `json:"interestedField" bson:"interestedField"`
	InterestedFieldDetail []string   `json:"interestedFieldDetail" bson:"interestedFieldDetail"`
	CreatedAt             time.Time  `json:"createdAt" bson:"createdAt"`
	UpdatedAt             time.Time  `json:"updatedAt" bson:"updatedAt"`
	DeletedAt             *time.Time `json:"deletedAt" bson:"deletedAt"`
}
