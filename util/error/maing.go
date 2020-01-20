package error

import "github.com/kyhsa93/gin-rest-example/util/error/http"

// Interface util error interface
type Interface interface {
	New() *Error
}

// Error error struct
type Error struct {
	HTTP http.Interface
}

// New create HTTP instance
func New() *Error {
	http := &http.HTTP{}
	return &Error{HTTP: http}
}
