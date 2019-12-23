package entity

// Account account entity for database table
type Account struct {
	Model
	Email    string
	Password string
}

// Accounts account entity list
type Accounts []Account
