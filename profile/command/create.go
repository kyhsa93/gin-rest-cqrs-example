package command

// CreateCommand create profile command
type CreateCommand struct {
	Email                 string
	AccountID             string
	Gender                string
	InterestedField       string
	InterestedFieldDetail []string
	FileID                string
}
