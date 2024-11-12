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

## Complete Example Walkthrough:

**Given the two sorted slices:**

- left = [2, 5, 8]
- right = [1, 3, 7, 9]

### The process proceeds as follows:

- First Iteration:

  - Compare 2 and 1 → 1 is smaller, so append 1.
  - result = [1]
  - Increment j (now j = 1).

- Second Iteration:

  - Compare 2 and 3 → 2 is smaller, so append 2.
  - result = [1, 2]
  - Increment i (now i = 1).

- Third Iteration:

  - Compare 5 and 3 → 3 is smaller, so append 3.
  - result = [1, 2, 3]
  - Increment j (now j = 2).

- Fourth Iteration:

  - Compare 5 and 7 → 5 is smaller, so append 5.
  - result = [1, 2, 3, 5]
  - Increment i (now i = 2).

- Fifth Iteration:

  - Compare 8 and 7 → 7 is smaller, so append 7.
  - result = [1, 2, 3, 5, 7]
  - Increment j (now j = 3).

- Sixth Iteration:
  - Compare 8 and 9 → 8 is smaller, so append 8.
  - result = [1, 2, 3, 5, 7, 8]
  - Increment i (now i = 3).
- Append Remaining Elements:
  - No elements left in left, but one element (9) left in right.
  - Append 9 to result.
  - result = [1, 2, 3, 5, 7, 8, 9]
