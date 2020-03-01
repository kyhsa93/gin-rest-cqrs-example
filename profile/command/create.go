package command

// CreateCommand create profile command
type CreateCommand struct {
	Email                 string
	AccountID             string
	Gender                string
	InterestedFeild       string
	InterestedFieldDetail []string
	FileID                string
}
