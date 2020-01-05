package dto

import "fmt"

// Account account dto for command action
type Account struct {
	Email    string `json:"email" example:"test@test.com"`
	SocialID string `json:"social_id" example:"socail_id"`
}

// Validate validate Account dto
func (account *Account) Validate(data *Account) error {
	if data.Email == "" {
		return fmt.Errorf("Can not validate email")
	}

	if data.SocialID == "" {
		return fmt.Errorf("Can not validate socialId")
	}

	return nil
}
