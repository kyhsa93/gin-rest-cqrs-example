package dto

import "time"

type Command struct {
	Name        string
	Description string
}

type Study struct {
	ID          string
	Name        string
	Description string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type Studies []Study
