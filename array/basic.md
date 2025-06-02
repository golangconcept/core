Here's a comprehensive overview of important **array and slice operations in Go**, with examples and use cases that are often useful in real-world coding and interviews.

---

## 📌 Arrays vs. Slices in Go

| Feature    | Array                    | Slice                                  |
| ---------- | ------------------------ | -------------------------------------- |
| Fixed Size | ✅ Yes                    | ❌ No (dynamic size)                    |
| Passed By  | Value (copy)             | Reference (points to underlying array) |
| Use Case   | Rare (system-level code) | Common in Go code                      |

---

## ✅ Essential Slice Operations & Functions

### 1. **Appending Elements**

```go
s := []int{1, 2}
s = append(s, 3, 4)
fmt.Println(s) // [1 2 3 4]
```

* Appending one or more elements.
* Can also append another slice: `append(s, otherSlice...)`

---

### 2. **Copying Slices**

```go
src := []int{1, 2, 3}
dst := make([]int, len(src))
copy(dst, src)
```

* `copy(dst, src)` copies **min(len(dst), len(src))** elements.
* `make` is used to initialize the destination.

---

### 3. **Deleting Elements**

#### a. Remove an element at index `i`:

```go
i := 2
s := []int{1, 2, 3, 4}
s = append(s[:i], s[i+1:]...)
fmt.Println(s) // [1 2 4]
```

---

### 4. **Insert an Element at Index**

```go
i := 2
s := []int{1, 2, 4, 5}
s = append(s[:i], append([]int{3}, s[i:]...)...)
fmt.Println(s) // [1 2 3 4 5]
```

---

### 5. **Reverse a Slice**

```go
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
```

---

### 6. **Check for Element Existence**

```go
func contains(s []int, target int) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}
```

---

### 7. **Sum of Elements**

```go
sum := 0
for _, v := range s {
	sum += v
}
```

---

### 8. **Filter Elements (e.g., remove even numbers)**

```go
func filterOdds(s []int) []int {
	result := []int{}
	for _, v := range s {
		if v%2 != 0 {
			result = append(result, v)
		}
	}
	return result
}
```

---

### 9. **Sorting a Slice**

```go
import "sort"

s := []int{3, 1, 2}
sort.Ints(s) // Ascending
sort.Sort(sort.Reverse(sort.IntSlice(s))) // Descending
```

---

### 10. **Multidimensional Slices**

```go
grid := make([][]int, 3)
for i := range grid {
	grid[i] = make([]int, 4)
}
```

---

## 🧠 Performance Notes

* Slices **share memory** with the original array—be cautious of modifying in-place.
* Use `copy` if you need a **true deep copy**.
* Slice operations are cheap, but appending a lot can cause reallocations.

---

## 🛠 Handy Built-in Functions

| Function              | Description                              |
| --------------------- | ---------------------------------------- |
| `len(s)`              | Length of slice                          |
| `cap(s)`              | Capacity (before reallocation is needed) |
| `append(a, b...)`     | Add elements or another slice            |
| `copy(dst, src)`      | Copy elements from `src` to `dst`        |
| `make([]T, len, cap)` | Create slice with length and capacity    |

---

Would you like examples of 2D slices, memory layout visuals, or benchmarking slice performance in Go?
