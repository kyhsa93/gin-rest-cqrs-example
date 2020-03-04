package command

// CreateCommand create account command
type CreateCommand struct {
	Email                 string
	Provider              string
	SocialID              string
	Password              string
	FCMToken              string
	Gender                string
	InterestedField       string
	InterestedFieldDetail []string
}
