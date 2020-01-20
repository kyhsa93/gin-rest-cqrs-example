package http

// Interface HTTP error interface
type Interface interface {
	Code() int
	Message() string
	BadRequest() *HTTP
	Unauthorized() *HTTP
	Forbidden() *HTTP
	NotFound() *HTTP
	Conflict() *HTTP
	InternalServerError() *HTTP
}

// HTTP http error struct
type HTTP struct {
	code    int
	message string
}

// Code return http error status code
func (http *HTTP) Code() int {
	return http.code
}

// Message return http error message
func (http *HTTP) Message() string {
	return http.message
}

// BadRequest return http 400 bad request error data
func (http *HTTP) BadRequest() *HTTP {
	return &HTTP{code: 400, message: "Bad request"}
}

// Unauthorized return http 401 unauthorized error data
func (http *HTTP) Unauthorized() *HTTP {
	return &HTTP{code: 401, message: "Unauthorized"}
}

// Forbidden return http 403 forbidden error data
func (http *HTTP) Forbidden() *HTTP {
	return &HTTP{code: 401, message: "Forbidden"}
}

// NotFound return http 404 not found error data
func (http *HTTP) NotFound() *HTTP {
	return &HTTP{code: 404, message: "Not found"}
}

// Conflict return http 209 conflict error data
func (http *HTTP) Conflict() *HTTP {
	return &HTTP{code: 409, message: "Conflict"}
}

// InternalServerError return http 500 internal server error data
func (http *HTTP) InternalServerError() *HTTP {
	return &HTTP{code: 500, message: "Internal server error"}
}
