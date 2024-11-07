# panic in go lang

`panic` and `recover` are mechanisms for handling runtime errors and controlling the flow of execution in the event of exceptional conditions.

They are part of Go's error handling model, but they differ from the standard `error` type handling and are typically used for situations where the program cannot continue execution safely.

```go
package main

import "fmt"

func riskyFunction() {
    fmt.Println("Something went wrong!")
    panic("Something terrible happened!")
}

func main() {
    fmt.Println("Program started.")
    riskyFunction() // This will panic and stop the program
    fmt.Println("This will not be executed.")
}
```
