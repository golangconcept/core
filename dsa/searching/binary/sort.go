package main

import (
	"fmt"
	"sort"
)

/*
sort package provides a convenient function sort.Search()
The time complexity of binary search is O(log n), where n is the number of elements in the slice.
*/

func BinarySearch(arr []int, target int) int {

	return sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})

}

func main1() {
	arr := []int{1, 5, 7, 12, 15, 16, 19, 20, 34}
	target := 7

	index := BinarySearch(arr, target)

	if index < len(arr) && arr[index] == target {
		fmt.Printf("The target is %d found at index %d\n", target, index)
	} else {
		fmt.Println("Target not found!!!")
	}
}
