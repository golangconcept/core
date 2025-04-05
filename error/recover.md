In Go, `recover` is used in conjunction with `defer` to handle panics. It allows you to regain control of a panicking goroutine and prevent the program from crashing. A panic typically causes the program to terminate, but `recover` can catch the panic and allow the program to continue running normally.

### Key Points about `recover`:
- `recover` only works when called inside a `defer` function.
- If there is no panic, `recover` returns `nil`.
- If a panic occurs and `recover` is invoked, it will stop the panic and return the value that was passed to `panic()`.
- After `recover` is called, the program will continue executing normally from the point where the panic was caught, skipping the remaining deferred functions.

### Example 1: **Basic Example with panic and recover**

```go
package main

import "fmt"

func riskyFunction() {
    panic("Something went wrong!") // This will cause a panic
}

func safeFunction() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic:", r)
        }
    }()

    riskyFunction() // This will trigger a panic
    fmt.Println("This will not be printed due to the panic")
}

func main() {
    safeFunction()
    fmt.Println("Program continues after recovery")
}
```

Output:
```
Recovered from panic: Something went wrong!
Program continues after recovery
```

### Explanation:
- The `riskyFunction()` causes a panic by calling `panic()`.
- Inside `safeFunction()`, we use `defer` to define an anonymous function that calls `recover()`.
- When the panic happens, the deferred function is executed, and `recover()` catches the panic.
- The program then continues running after recovering from the panic, and the message `"Program continues after recovery"` is printed.

### Example 2: **Using recover to prevent a crash from multiple panics**

```go
package main

import "fmt"

func causePanic1() {
    panic("First panic!")
}

func causePanic2() {
    panic("Second panic!")
}

func handlePanics() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from:", r)
        }
    }()
    
    causePanic1() // This will panic
    causePanic2() // This will also panic, but we can recover from it
}

func main() {
    handlePanics()
    fmt.Println("Program continues without crashing")
}
```

Output:
```
Recovered from: First panic!
Recovered from: Second panic!
Program continues without crashing
```

### Explanation:
- The `handlePanics()` function handles multiple panics.
- Even though both `causePanic1()` and `causePanic2()` trigger panics, we can recover from them and allow the program to continue running.
- `recover()` is called after each panic, and it prevents the program from crashing.

### Example 3: **Recover with custom error handling**

```go
package main

import "fmt"

func riskyOperation() {
    panic("Critical error occurred!")
}

func safeOperation() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println("Handled error:", err)
        }
    }()
    riskyOperation() // This will trigger a panic
    fmt.Println("This won't be printed because of panic")
}

func main() {
    fmt.Println("Before the safe operation")
    safeOperation()
    fmt.Println("After the safe operation")
}
```

Output:
```
Before the safe operation
Handled error: Critical error occurred!
After the safe operation
```

### Explanation:
- The `riskyOperation()` function triggers a panic.
- In the `safeOperation()` function, we catch the panic using `recover()` and handle it by printing a custom error message.
- This prevents the program from crashing, and the program continues after the panic is handled.

### Example 4: **Recover and Return to Normal Flow**

```go
package main

import "fmt"

func riskyTask() string {
    panic("Something unexpected happened!")
}

func performSafeTask() (result string) {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recovered from panic in performSafeTask:", r)
            result = "Task failed, but recovered"
        }
    }()
    
    result = riskyTask() // This triggers a panic
    return result
}

func main() {
    res := performSafeTask()
    fmt.Println("Result of task:", res)
}
```

Output:
```
Recovered from panic in performSafeTask: Something unexpected happened!
Result of task: Task failed, but recovered
```

### Explanation:
- In this example, `riskyTask()` triggers a panic.
- The `defer` function in `performSafeTask()` catches the panic, allowing the function to return a default value (`"Task failed, but recovered"`) instead of crashing.
- The program continues without interruption after the panic is handled.

### When to use `recover`:
- **Error handling**: `recover` is useful when you want to handle unexpected errors (such as panics) in a controlled way.
- **Graceful recovery**: In situations where you don’t want your program to crash but need to stop execution of a function due to a panic (e.g., an out-of-bounds error or other unexpected conditions).
- **Cleanup in a panic scenario**: If a function uses `defer` to clean up resources, `recover` can ensure that the cleanup still happens even after a panic.

### Important Notes:
- **`recover` only works in deferred functions**: If you call `recover` outside of a deferred function, it will return `nil` and have no effect.
- **Recovering from panics should be done cautiously**: While `recover` is useful for preventing crashes, overuse of panic/recover for error handling is generally discouraged. It’s better to use explicit error handling (returning `error` types) for expected conditions.

Would you like to see more examples, or do you have a specific scenario in mind that you’d like to explore further?