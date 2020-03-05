package body

// CreateProfile request body for POST /profiles
type CreateProfile struct {
	AccountID             string   `json:"accountId" example:"accountId"`
	Email                 string   `json:"email" example:"test@gmail.com"`
	Gender                string   `json:"gender" example:"male"`
	InterestedField       string   `json:"interestedField" example:"develop"`
	InterestedFieldDetail []string `json:"interestedFieldDetail" example:"web,server"`
}

// UpdateProfile request body for PUT /profiles
type UpdateProfile struct {
	InterestedField       string   `json:"interestedField" example:"develop"`
	InterestedFieldDetail []string `json:"interestedFieldDetail" example:"web,server"`
	FileID                string   `json:"fileId" example:"fileId"`
}
