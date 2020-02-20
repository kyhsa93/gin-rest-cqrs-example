package command

// UpdateCommand update account command
type UpdateCommand struct {
	AccountID string
	Email     string
	Provider  string
	SocialID  string
	Password  string
	ImageKey  string
	Gender    string
	Interest  string
}
