package entity

import "time"

// File file entity
type File struct {
	ID        string    `json:"_id" bson:"_id"`
	AccountID string    `json:"accountId" bson:"accountId"`
	Usage     string    `json:"usage" bson:"usage"`
	CreatedAt time.Time `json:"createdAt" bson:"createdAt"`
}
