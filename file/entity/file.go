package entity

// File file entity
type File struct {
	Model
	AccountID string `gorm:"not null"`
	Usage     string `gorm:"not null"`
}
