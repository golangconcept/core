### Explanation:

**merge function:**

- This function takes two sorted slices (left and right) and merges them into a single sorted slice.
- It compares the elements of left and right, appending the smaller element to the result slice.
- After the loop, any remaining elements from either slice are appended to the result.

### mergeSort function:

- This is the recursive function that splits the array until the base case is reached (arrays with 1 or 0 elements are trivially sorted).
- It divides the array into two halves and calls itself recursively to sort each half.
- Finally, the merge function is called to combine the two sorted halves into a single sorted array.

### Merge function merges two sorted slices into a single sorted slice

```go
func merge(left, right []int) []int {
result := []int{}
i, j := 0, 0

    // Compare elements from both slices and append the smaller one to the result slice
    for i < len(left) && j < len(right) {
    	if left[i] <= right[j] {
    		result = append(result, left[i])
    		i++
    	} else {
    		result = append(result, right[j])
    		j++
    	}
    }

    // Append any remaining elements from both slices
    result = append(result, left[i:]...)
    result = append(result, right[j:]...)

    return result

}
```

### MergeSort function recursively splits the array and merges them

```go
// MergeSort function recursively splits the array and merges them
func mergeSort(arr []int) []int {
	// Base case: if the array has 1 or 0 elements, it is already sorted
	if len(arr) <= 1 {
		return arr
	}

	// Split the array into two halves
	mid := len(arr) / 2
	left := mergeSort(arr[:mid])  // Recursively sort the left half
	right := mergeSort(arr[mid:]) // Recursively sort the right half

	// Merge the sorted halves
	return merge(left, right)
}
```
