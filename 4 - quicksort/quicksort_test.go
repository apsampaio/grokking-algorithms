package main

import "testing"

func TestQuicksort(t *testing.T) {
	arr := []int{5, 1, 2, 10, 4, 3}

	r := Quicksort(arr)

	if r[0] != 1 {
		t.Errorf("expected %d at index 0 and got %d", 1, r[0])
	}

	if r[len(r)-1] != 10 {
		t.Errorf("expected %d at last index and got %d", 10, r[len(r)-1])
	}
}
