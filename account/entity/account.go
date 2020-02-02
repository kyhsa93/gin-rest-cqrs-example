package entity

// Account account entity for database table
type Account struct {
	Model
	Email    string `gorm:"unique;not null"`
	SocialID string `gorm:"not null"`
	Password string
	Provider string `gorm:"not null"`
	ImageKey string
	Gender   string `gorm:"not null"`
}
