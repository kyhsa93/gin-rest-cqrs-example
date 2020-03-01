package dto

// Profile profile dto
type Profile struct {
	AccountID             string   `json:"accountId" example:"account_id"`
	Email                 string   `json:"email" example:"test@gmail.com"`
	Gender                string   `json:"gender" example:"male"`
	InterestedField       string   `json:"interestedField" example:"develop"`
	InterestedFieldDetail []string `json:"interestedFieldDetail" example:"web,server"`
	FileID                string   `json:"fileId" example:"fileId"`
}
