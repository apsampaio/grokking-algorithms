package main

import "testing"

func TestSelectionSort(t *testing.T) {
	arr := []int{5, 3, 6, 2, 10}

	r := SelectionSort(arr)

	if r[0] != 2 {
		t.Errorf("expected %d and received %d", 2, r[0])
	}

	if r[len(arr)-1] != 10 {
		t.Errorf("expected %d and received %d", 10, r[len(arr)-1])
	}
}
