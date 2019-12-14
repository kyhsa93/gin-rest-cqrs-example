package model

import "time"

type Command struct {
}

type Account struct {
	ID        string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Accounts []Account
