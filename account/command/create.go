package command

// CreateCommand create account command
type CreateCommand struct {
	Email                 string
	Provider              string
	SocialID              string
	Password              string
	Gender                string
	InterestedField       string
	InterestedFieldDetail []string
}
