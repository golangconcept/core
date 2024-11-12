package main

import "fmt"

func main() {
	arr := []int{9, 7, 5, 11, 12, 2, 14, 3, 10, 6}
	fmt.Println("Original array:", arr)

	// Apply Quick Sort
	quick(arr)

	// Output the sorted array
	fmt.Println("Sorted array:", arr)
}
func partition(arr []int) int {
	i := -1
	pivot := arr[len(arr)-1]

	for j := 0; j < len(arr)-1; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[len(arr)-1] = arr[len(arr)-1], arr[i+1]
	return i + 1
}

func quick(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivotIndex := partition(arr)

	quick(arr[:pivotIndex])
	quick(arr[pivotIndex+1:])
}
