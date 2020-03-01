package dto

// Profile profile dto
type Profile struct {
	AccountID             string   `json:"accountId" example:"account_id"`
	Email                 string   `json:"email" example:"test@gmail.com"`
	Gender                string   `json:"gender" example:"male"`
	InterestedFeild       string   `json:"interested_field" example:"develop"`
	InterestedFieldDetail []string `json:"interested_field_detail" example:"web,server"`
	FileID                string   `json:"file_id" example:"file_id"`
}
