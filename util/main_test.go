package util_test

import (
	"testing"

	"github.com/kyhsa93/gin-rest-cqrs-example/util"
)

// TestInitialize test InitizlizeUtil method
func TestInitializeUtil(t *testing.T) {
	instance := util.Initialize()
	if instance == nil {
		t.Error("Can not create util instance")
	}
}
