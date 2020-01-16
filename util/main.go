package util

import (
	"github.com/kyhsa93/gin-rest-example/util/error"
)

// Util provide utilities
type Util struct {
	Error *error.Error
}

// InitializeUtil initialize utilities
func InitializeUtil() *Util {
	error := error.New()
	return &Util{Error: error}
}
