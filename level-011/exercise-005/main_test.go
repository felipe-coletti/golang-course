package main

import "testing"

func TestDouble(t *testing.T) {
	result := double(5)

	expected := 10

	if result != expected {
		t.Errorf("double(5) = %d; esperado %d", result, expected)
	}
}
