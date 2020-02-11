package http_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-cqrs-example/util/error/http"
)

// TestCode test test http error method named Code
func TestCode(t *testing.T) {
	instance := &http.HTTP{}
	code := instance.Code()
	if code != 0 {
		t.Error("Util HTTP Error Code method error")
	}
}

// TestMessage test test http error method named Message
func TestMessage(t *testing.T) {
	instance := &http.HTTP{}
	code := instance.Message()
	if code != "" {
		t.Error("Util HTTP Error Message method error")
	}
}

// TestBadRequest test test http error method named Message
func TestBadRequest(t *testing.T) {
	instance := &http.HTTP{}
	httpError := instance.BadRequest()
	if httpError == nil {
		t.Error("Util HTTP Error TestBadRequest method error")
	}
}

// TestUnauthorized test test http error method named Message
func TestUnauthorized(t *testing.T) {
	instance := &http.HTTP{}
	httpError := instance.Unauthorized()
	if httpError == nil {
		t.Error("Util HTTP Error TestUnauthorized method error")
	}
}

// TestForbidden test test http error method named Message
func TestForbidden(t *testing.T) {
	instance := &http.HTTP{}
	httpError := instance.Forbidden()
	if httpError == nil {
		t.Error("Util HTTP Error TestForbidden method error")
	}
}

// TestNotFound test test http error method named Message
func TestNotFound(t *testing.T) {
	instance := &http.HTTP{}
	httpError := instance.NotFound()
	if httpError == nil {
		t.Error("Util HTTP Error TestNotFound method error")
	}
}

// TestConflict test test http error method named Message
func TestConflict(t *testing.T) {
	instance := &http.HTTP{}
	httpError := instance.Conflict()
	if httpError == nil {
		t.Error("Util HTTP Error TestConflict method error")
	}
}

// TestInternalServerError test test http error method named Message
func TestInternalServerError(t *testing.T) {
	instance := &http.HTTP{}
	httpError := instance.InternalServerError()
	if httpError == nil {
		t.Error("Util HTTP Error TestInternalServerError method error")
	}
}
