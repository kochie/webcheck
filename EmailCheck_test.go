package webcheck

import (
	"testing"
)

func TestIsValidEmail(t *testing.T){
	email := "hello@example.com"
	isValid := IsValidEmail(email)
	if isValid != true {
		t.Error("Expected true got ", isValid)
	}
}
