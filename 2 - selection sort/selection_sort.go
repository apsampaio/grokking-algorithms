package main

func SelectionSort(arr []int) []int {
	newArr := []int{}
	size := len(arr)

	for i := 0; i < size; i++ {
		idx := findSmallest(arr)
		newArr = append(newArr, arr[idx])
		arr = append(arr[:idx], arr[idx+1:]...)
	}

	return newArr
}

func findSmallest(arr []int) int {
	s := arr[0]
	idx := 0

	for i := 1; i < len(arr); i++ {
		if arr[i] < s {
			s = arr[i]
			idx = i
		}
	}

	return idx
}
