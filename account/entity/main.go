package entity

import (
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql" // mysql package for gorm
)

// Model default entity struct
type Model struct {
	ID        string `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
