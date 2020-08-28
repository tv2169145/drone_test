package main

import "testing"

func TestGetHelloWorld(t *testing.T) {
	result := GetHelloWorld()
	if result != "hello jimmy 1234" {
		t.Fatal("should get hello jimmy")
	}
}
