package webcheck

import (
	"testing"
	"fmt"
)

func TestIsValidEmail(t *testing.T){
	emailList := []string{
		"hello@example.com",
		"Abc\\@def@example.com",
		"Fred\\ Bloggs@example.com",
		"Joe.\\\\Blow@example.com",
		"customer/department=shipping@example.com",
		"$A12345@example.com",
		"!def!xyz%abc@example.com",
		"_somename@example.com",
	}
	for _, email := range emailList {
	isValid := IsValidEmail(email)
	if isValid != true {
		fmt.Println("Email couldn't be validated ", email)
		t.Error("Expected true got ", isValid)
	}
}

	emailList = []string{
		"\"Abc@def\"@example.com",
		"\"Fred Bloggs\"@example.com",
	}
	for _, email := range emailList {
	isValid := IsValidEmail(email)
	if isValid != false {
		fmt.Println("Email shouldn't be validated ", email)
		t.Error("Expected true got ", isValid)
	}
}
}
