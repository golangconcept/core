Example Walkthrough of the Partition Function
Let’s walk through the partition function with a simple example:

**Array:** [9, 7, 5, 11, 12, 2, 14, 3, 10, 6]
**Pivot: 6** (last element)

## Initial state:

- i = -1
- pivot = 6
- arr = [9, 7, 5, 11, 12, 2, 14, 3, 10, 6]

### Scan through the array with j:

- j = 0: arr[j] = 9 → 9 > 6 → do nothing.
- j = 1: arr[j] = 7 → 7 > 6 → do nothing.
- j = 2: arr[j] = 5 → 5 <= 6 → increment i to 0, swap arr[0] and arr[2].

#### Array becomes: [5, 7, 9, 11, 12, 2, 14, 3, 10, 6]

- j = 3: arr[j] = 11 → 11 > 6 → do nothing.
- j = 4: arr[j] = 12 → 12 > 6 → do nothing.
- j = 5: arr[j] = 2 → 2 <= 6 → increment i to 1, swap arr[1] and arr[5].

#### Array becomes: [5, 2, 9, 11, 12, 7, 14, 3, 10, 6]

- j = 6: arr[j] = 14 → 14 > 6 → do nothing.
- j = 7: arr[j] = 3 → 3 <= 6 → increment i to 2, swap arr[2] and arr[7].

#### Array becomes: [5, 2, 3, 11, 12, 7, 14, 9, 10, 6]

- j = 8: arr[j] = 10 → 10 > 6 → do nothing.

### Final swap:

After finishing the loop, we swap the pivot 6 with `arr[i + 1]`, i.e., swap arr[3] with arr[9].

#### Array becomes: [5, 2, 3, 6, 12, 7, 14, 9, 10, 11]

The pivot (6) is now in its correct sorted position at index 3.
