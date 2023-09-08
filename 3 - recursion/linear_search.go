package main

import "errors"

func LinearSearch(arr []int, value int, idx int) (int, error) {
	if idx == len(arr) {
		return -1, errors.New("value not found")
	}
	if arr[idx] == value {
		return idx, nil
	}
	idx += 1
	return LinearSearch(arr, value, idx)
}
