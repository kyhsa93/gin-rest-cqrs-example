package dto

// Account account dto for command action
type Account struct {
	Email           string `json:"email" example:"test@test.com"`
	SocialID        string `json:"social_id" example:"socail_id"`
	ProdileImageURL string `json:"profile_image_url" example:"profile_image_url.com"`
}
