package dto

// Account account dto for command action
type Account struct {
	Email    string `json:"email" example:"test@test.com"`
	Password string `json:"password" example:"testpassword"`
}
