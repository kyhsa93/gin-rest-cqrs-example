package util

import (
	"github.com/kyhsa93/gin-rest-cqrs-example/util/error"
)

// Interface Utile interface
type Interface interface {
}

// Util provide utilities
type Util struct {
	Error *error.Error
}

// Initialize initialize utilities
func Initialize() *Util {
	error := error.New()
	return &Util{Error: error}
}
