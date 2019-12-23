package model

import "time"

// Account account model
type Account struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// Accounts account model list
type Accounts []Account
