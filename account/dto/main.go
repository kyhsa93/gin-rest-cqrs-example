package dto

// Account account dto for command action
type Account struct {
	Email    string `json:"email" example:"test@test.com" binding:"required"`
	SocialID string `json:"social_id" example:"social_id" binding:"required"`
}
