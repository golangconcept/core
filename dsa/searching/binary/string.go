package main

import (
	"fmt"
	"sort"
)

func BinarySearchString(arr []string, target string) int {
	return sort.Search(len(arr), func(i int) bool {
		return arr[i] >= target
	})
}

func main() {
	stringsArr := []string{"apple", "banana", "cherry", "date", "fig"}
	target := "cherry"

	// Find index using sort.Search
	index := BinarySearchString(stringsArr, target)

	// Check if the element exists
	if index < len(stringsArr) && stringsArr[index] == target {
		fmt.Printf("Target '%s' found at index %d\n", target, index)
	} else {
		fmt.Println("Target not found")
	}
}
