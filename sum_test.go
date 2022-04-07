package main

import "testing"

func TestSum(t *testing.T) {
	expected := sum(1, 2)
	result := 3
	if expected != result {
		t.Errorf("Salah hasil, harusnya %v", result)
	}
}
