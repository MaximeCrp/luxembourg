package main

import "testing"

func TestMultiply(t *testing.T) {
	expected := multiply(6, 7)

	if expected != 42 {
		t.Fatal("should be 42")
	}
}
