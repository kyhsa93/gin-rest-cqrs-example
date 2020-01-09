package error

// Error error struct
type Error struct {
	HTTP *HTTP
}

// New create HTTP instance
func New() *Error {
	http := &HTTP{}
	return &Error{HTTP: http}
}
