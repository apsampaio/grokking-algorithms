package main

func Quicksort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := []int{}
	greater := []int{}

	for i := 1; i < len(arr); i++ {
		if arr[i] <= pivot {
			less = append(less, arr[i])
		} else {
			greater = append(greater, arr[i])
		}
	}

	return append(append(Quicksort(less), pivot), Quicksort(greater)...)
}
