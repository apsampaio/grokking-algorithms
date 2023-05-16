package main

import "testing"

func TestBinarySearch(t *testing.T) {

	arr := []int{1, 3, 5, 7, 9}
	r := BinarySearch(arr, 3)

	if r != 1 {
		t.Errorf("expected  1  received %d", r)
	}

	r = BinarySearch(arr, -1)

	if r != -1 {
		t.Errorf("expected -1 received %d", r)
	}
}
