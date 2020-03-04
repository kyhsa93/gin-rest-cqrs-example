package body

// CreateAccount request body for POST /accounts
type CreateAccount struct {
	Email                 string   `json:"email" example:"test@gmail.com" binding:"required"`
	Provider              string   `json:"provider" example:"gmail" binding:"required"`
	SocialID              string   `json:"socialId" example:"socialId"`
	Password              string   `json:"password" example:"password"`
	FCMToken              string   `json:"fcmToken" example:"fcmToken"`
	Gender                string   `json:"gender" example:"male"`
	InterestedField       string   `json:"interestedField" example:"develop"`
	InterestedFieldDetail []string `json:"interestedFieldDetail" example:"web,server"`
}

// UpdateAccount request body for PUT /accounts
type UpdateAccount struct {
	Password              string   `json:"password" example:"password"`
	FCMToken              string   `json:"fcmToken" example:"fcmToken"`
	InterestedField       string   `json:"interestedField" example:"develop"`
	InterestedFieldDetail []string `json:"interestedFieldDetail" example:"web,server"`
}
