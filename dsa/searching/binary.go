package main

import "fmt"

/*
`Binary` search is a classic algorithm used to search for an element in a sorted collection
(e.g., an array or slice) by repeatedly dividing the search interval in half.
*/

/*
Time complexity of O(log n), making it ideal for searching in sorted datasets.
*/

func BinarySearch(arr []int, target int) int {

	left, right := 0, len(arr)-1

	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1 // target not found
}
func main() {

	arr := []int{1, 3, 5, 6, 8, 10, 13, 30}
	target := 8

	index := BinarySearch(arr, target)

	if index != -1 {
		fmt.Printf("Target %d found at index: %d\n", target, index)

	} else {
		fmt.Printf("Target not found")
	}

}
