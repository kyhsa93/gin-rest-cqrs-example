package model

import "time"

// Account account model
type Account struct {
	ID        string    `json:"id" example:"389df385-ccaa-49c1-aee2-698ba1191857"`
	Email     string    `json:"email" example:"test@test.com"`
	CreatedAt time.Time `json:"created_at" example:"2019-12-23 12:27:37"`
	UpdatedAt time.Time `json:"updated_at" example:"2019-12-23 12:27:37"`
}
