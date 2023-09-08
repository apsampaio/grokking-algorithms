package main

import "testing"

func TestLinearSearch(t *testing.T) {
	arr := []int{5, 3, 6, 2, 10}

	r, _ := LinearSearch(arr, 6, 0)

	if r != 2 {
		t.Errorf("expected %d and received %d", 2, r)
	}

	r, err := LinearSearch(arr, 42, 0)

	if err == nil {
		t.Errorf("expected error and received %d", r)
	}
}
