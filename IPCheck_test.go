package webcheck

import (
	"testing"
	"fmt"
)

func TestIsValidIP(t *testing.T){
	IPList := []string{
		"255.255.255.255",
		"192.168.17.1",
		"192.0.0.1",
		"192.168.17.230",
		"0.0.0.0",
		"2001:0db8:0000:0000:0000:ff00:0042:8329",
	}
	for _, ip := range IPList {
		isValid := IsValidIP(ip)
		if isValid != true {
			fmt.Println("IP couldn't be validated ", ip)
			t.Error("Expected true got ", isValid)
		}
	}

	IPList = []string{
		"256.12.214.32",
		"0..0.1",
	}
	for _, ip := range IPList {
		isValid := IsValidIP(ip)
		if isValid != false {
			fmt.Println("IP shouldn't be validated ", ip)
			t.Error("Expected false got ", isValid)
		}
	}
}

func TestIsValidIPv6(t *testing.T) {
	IPList := []string{
		"2001:0db8:0000:0000:0000:ff00:0042:8329",
		"192.168.17.1",
		"192.0.0.1",
		"192.168.17.230",
		"",
	}
	for _, ip := range IPList {
	isValid := IsValidIPv6(ip)
	if isValid != true {
		fmt.Println("IP couldn't be validated ", ip)
		t.Error("Expected true got ", isValid)
	}
}

	IPList = []string{
		"256.12.214.32",
		"0..0.1",
	}
	for _, ip := range IPList {
		isValid := IsValidIPv6(ip)
		if isValid != false {
			fmt.Println("IP shouldn't be validated ", ip)
			t.Error("Expected false got ", isValid)
		}
	}
	}

