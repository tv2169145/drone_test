package utils

import "testing"

func TestGetHelloWorld(t *testing.T) {
	result := GetHelloWorld()
	if result != "hello jimmy" {
		t.Error("should get jimmy")
	}
}
