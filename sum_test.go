package main

import "testing"

func TestSum(t *testing.T) {
	results := sum(2,3)

	if results != 5 {
		t.Error("The result must be 5")
	}
}