package repositories

import (
	"study/model"

	"github.com/jinzhu/gorm"
)

type StudyRepository struct {
	db *gorm.DB
}

func NewStudyRepository(db *gorm.DB) *StudyRepository {
	return &StudyRepository{db: db}
}

func (repository *StudyRepository) Save(study *model.Study) {
	repository.db.Save(study)
}

func (repository *StudyRepository) FindAll() model.Studies {
	var studies model.Studies
	repository.db.Find(&studies)
	return studies
}

func (repository *StudyRepository) FindById(id string) {
	var study model.Study
	repository.db.Take(&study).Where(&model.Model{ID: id})
}

func (repository *StudyRepository) Delete(id string) {
	repository.db.Delete(&model.Model{ID: id})
}
