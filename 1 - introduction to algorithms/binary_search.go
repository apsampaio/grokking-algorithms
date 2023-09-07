package main

// 1.1 Suppose you have a sorted list of 128 names, and you’re searching
// through it using binary search. What’s the maximum number of
// steps it would take?
// Log2(128) => 7

// 1.2 Suppose you double the size of the list. What’s the maximum
// number of steps now?
// Log2(256) => 8

// Receives a sorted array to execute a binary search.
// returns the index from the target, -1 in case not found.
func BinarySearch(arr []int, item int) int {
	var mid, guess, low int
	high := len(arr) - 1

	for {
		// check middle element
		mid = (low + high) / 2
		guess = arr[mid]

		// found item
		if guess == item {
			return mid
		}

		if guess > item {
			high = mid - 1 // guess too high
		} else {
			low = mid + 1 // guess too low
		}

		// break loop
		if low > high {
			break
		}
	}

	return -1
}
