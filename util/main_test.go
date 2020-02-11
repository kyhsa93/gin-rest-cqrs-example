package util_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

// TestInitializeUtil test InitizlizeUtil method
func TestInitializeUtil(t *testing.T) {
	instance := util.InitializeUtil()
	if instance == nil {
		t.Error("Can not create util instance")
	}
}
