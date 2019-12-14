package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Model struct {
	ID        string `gorm:"primary_key;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}

func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	uuid, _ := uuid.NewRandom()
	return scope.SetColumn("ID", uuid.String())
}
