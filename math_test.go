package main

import "testing"

func TestSum(t *testing.T) {
	if sum(1, 2) != 3 {
		t.Errorf("Sum(1, 2) = %d; want 3", sum(1, 2))
	}
}