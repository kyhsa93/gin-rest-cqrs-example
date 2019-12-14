package repository

import (
	"github.com/kyhsa93/go-rest-example/study/dto"
	"github.com/kyhsa93/go-rest-example/study/entity"

	"github.com/jinzhu/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (repository *Repository) Save(data *dto.Study) {
	study := &entity.Study{Name: data.Name, Description: data.Description}
	err := repository.db.Save(study).Error
	if err != nil {
		panic(err)
	}
}

func (repository *Repository) FindAll() entity.Studies {
	var studies entity.Studies
	err := repository.db.Find(&studies).Error
	if err != nil {
		panic(err)
	}
	return studies
}

func (repository *Repository) FindById(id string) entity.Study {
	var study entity.Study
	repository.db.Where(&entity.Study{Model: entity.Model{ID: id}}).Take(&study)
	return study
}

func (repository *Repository) Delete(id string) {
	err := repository.db.Delete(&entity.Study{Model: entity.Model{ID: id}}).Error
	if err != nil {
		panic(err)
	}
}
