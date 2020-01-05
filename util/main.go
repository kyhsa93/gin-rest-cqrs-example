package util

import (
	httpError "github.com/kyhsa93/go-rest-example/util/http-error"
)

// Util provide utilities
type Util struct {
	HTTPError *httpError.HTTPError
}

// InitializeUtil initialize utilities
func InitializeUtil() *Util {
	httpError := httpError.New()
	return &Util{HTTPError: httpError}
}
