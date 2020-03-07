package command

// UpdateProfileCommand update profile command
type UpdateProfileCommand struct {
	ID                    string
	InterestedField       string
	InterestedFieldDetail []string
	FileID                string
}
