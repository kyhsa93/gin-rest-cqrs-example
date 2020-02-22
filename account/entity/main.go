package entity

import (
	"time"
)

// Model default entity struct
type Model struct {
	ID        string `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
