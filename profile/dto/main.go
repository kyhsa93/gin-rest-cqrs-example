package dto

// CreateProfile dto for create profile
type CreateProfile struct {
	AccountID             string
	Email                 string
	Gender                string
	InterestedField       string
	InterestedFieldDetail []string
}

// UpdateProfile request body for PUT /profiles
type UpdateProfile struct {
	AccountID             string
	InterestedField       string
	InterestedFieldDetail []string
	FileID                string
}
