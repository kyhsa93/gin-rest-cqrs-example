package entity

// Account account entity for database table
type Account struct {
	Model
	Email    string `gorm:"unique;not null"`
	SocialID string `gorm:"not null"`
}

// Accounts account entity list
type Accounts []Account
