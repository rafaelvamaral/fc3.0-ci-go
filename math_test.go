package main

import "testing"

func TestSoma(t *testing.T) {
	result := Soma(10, 15)
	expected := 25

	if result != expected {
		t.Errorf("Soma(10, 15) = %d; want %d", result, expected)
	}
}
