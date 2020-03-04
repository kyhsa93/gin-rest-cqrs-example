package command

// UpdateCommand update account command
type UpdateCommand struct {
	AccountID             string
	Password              string
	FCMToken              string
	InterestedField       string
	InterestedFieldDetail []string
}
