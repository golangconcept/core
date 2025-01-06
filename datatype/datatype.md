# Data Types

The following are the basic data types available in Go

- bool
- Numeric Types
  - int8, int16, int32, int64, int
  - uint8, uint16, uint32, uint64, uint
  - float32, float64
  - complex64, complex128
  - byte
  - rune
- string

**There are also `byte` and `rune`. They are equivalent to `uint8` and `int32`, respectively.**

Note: In computer programming, `monkey patching` is a technique used to dynamically update the behavior of a piece of code at run-time. `It is used to extend or modify the runtime code of dynamic languages` such as `Smalltalk`, JavaScript, `Objective-C`, `Ruby`, `Perl`, `Python`, `Groovy`, and `Lisp` without altering the original source code.

### Basic types

![basic data types](image.png)

- The variable i has type int, represented in memory as a single 32-bit word. (All these pictures show a 32-bit memory layout; in the current implementations, only the pointer gets bigger on a 64-bit machine—int is still 32 bits—though an implementation could choose to use 64 bits instead.)

- The variable j has type int32, because of the explicit conversion. Even though i and j have the same memory layout, they have different types: the assignment i = j is a type error and must be written with an explicit conversion: i = int(j).

- The variable f has type float, which the current implementations represent as a 32-bit floating-point value. It has the same memory footprint as the int32 but a different internal layout.

#### If you need to compare types you defined, you shouldn't use `reflect.TypeOf(xxx)`. Instead, use `reflect.TypeOf(xxx).Kind()`.

**There are two categories of types:**

- direct types (the types you defined directly)
- basic types (`int`, `float64`, `struct`, ...)

> Here's the conclusion.

If you need to compare with basic types, use `reflect.TypeOf(xxx)`.Kind(); and if you need to compare with self-defined types, use `reflect.TypeOf(xxx)`.

````go
package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	j := 3.4
	k := j
	l := int32(4)
	fmt.Println(j, k, l, unsafe.Sizeof(l), reflect.TypeOf(l), reflect.TypeOf(l).Kind())
}
````

In **Go (Golang)**, types are a fundamental part of the language. Go is a statically typed language, which means variables must be declared with a type, and type safety is enforced at compile time.

Here is a breakdown of the **different types in Go**, with examples for each:

---

### **1. Basic Types**

These are the standard primitive types used to represent values.

#### **a. Integer Types**
Go has several integer types, both signed and unsigned, and with various bit-widths.

- **int**: A signed integer (typically 32-bit or 64-bit, depending on the architecture).
- **int8**: A signed 8-bit integer.
- **int16**: A signed 16-bit integer.
- **int32**: A signed 32-bit integer (commonly used for Unicode characters, `rune` type).
- **int64**: A signed 64-bit integer.
- **uint**: An unsigned integer (typically 32-bit or 64-bit).
- **uint8**: An unsigned 8-bit integer (same as `byte`).
- **uint16**: An unsigned 16-bit integer.
- **uint32**: An unsigned 32-bit integer.
- **uint64**: An unsigned 64-bit integer.

**Example:**

```go
package main

import "fmt"

func main() {
    var x int = 42
    var y uint8 = 255
    var z int64 = 1234567890

    fmt.Println(x)  // 42
    fmt.Println(y)  // 255
    fmt.Println(z)  // 1234567890
}
```

#### **b. Floating-Point Types**
Go has two floating-point types: `float32` and `float64`.

- **float32**: A 32-bit floating-point number.
- **float64**: A 64-bit floating-point number (default for floating-point numbers).

**Example:**

```go
package main

import "fmt"

func main() {
    var a float32 = 3.14
    var b float64 = 3.14159

    fmt.Println(a)  // 3.14
    fmt.Println(b)  // 3.14159
}
```

#### **c. Boolean Type**
- **bool**: Represents a boolean value, which can either be `true` or `false`.

**Example:**

```go
package main

import "fmt"

func main() {
    var flag bool = true
    fmt.Println(flag)  // true
}
```

#### **d. String Type**
- **string**: Represents a sequence of Unicode characters.

**Example:**

```go
package main

import "fmt"

func main() {
    var name string = "Hello, Go!"
    fmt.Println(name)  // Hello, Go!
}
```

---

### **2. Composite Types**

Composite types are more complex types built from basic types.

#### **a. Arrays**
Arrays are fixed-size collections of elements of the same type.

**Example:**

```go
package main

import "fmt"

func main() {
    var arr [3]int = [3]int{1, 2, 3}
    fmt.Println(arr)  // [1 2 3]
}
```

#### **b. Slices**
Slices are more flexible than arrays and can grow and shrink dynamically. They are a reference type.

**Example:**

```go
package main

import "fmt"

func main() {
    var slice []int = []int{1, 2, 3, 4}
    fmt.Println(slice)  // [1 2 3 4]

    // Add an element to the slice
    slice = append(slice, 5)
    fmt.Println(slice)  // [1 2 3 4 5]
}
```

#### **c. Maps**
Maps are unordered collections of key-value pairs.

**Example:**

```go
package main

import "fmt"

func main() {
    var m map[string]int = map[string]int{"Alice": 25, "Bob": 30}
    fmt.Println(m)  // map[Alice:25 Bob:30]

    // Access a value using a key
    fmt.Println(m["Alice"])  // 25
}
```

#### **d. Structs**
Structs are used to group different types of data together.

**Example:**

```go
package main

import "fmt"

// Define a struct
type Person struct {
    Name string
    Age  int
}

func main() {
    // Create an instance of the struct
    var p Person = Person{Name: "Alice", Age: 25}
    fmt.Println(p)  // {Alice 25}

    // Access struct fields
    fmt.Println(p.Name)  // Alice
    fmt.Println(p.Age)   // 25
}
```

#### **e. Pointers**
Pointers store the memory address of a variable.

**Example:**

```go
package main

import "fmt"

func main() {
    var x int = 58
    var ptr *int = &x  // ptr points to the memory address of x
    fmt.Println(ptr)   // Memory address of x
    fmt.Println(*ptr)  // 58, dereferencing the pointer to get the value
}
```

---

### **3. Interface Types**

Interfaces allow you to define methods that can be implemented by any type, enabling polymorphism.

**Example:**

```go
package main

import "fmt"

// Define an interface
type Speaker interface {
    Speak() string
}

// Define a struct
type Person struct {
    Name string
}

// Implement the Speak method for Person type
func (p Person) Speak() string {
    return "Hello, my name is " + p.Name
}

func main() {
    var speaker Speaker = Person{Name: "Alice"}
    fmt.Println(speaker.Speak())  // Hello, my name is Alice
}
```

---

### **4. Type Aliases and Type Conversions**

You can define an alias for an existing type or convert between different types.

#### **a. Type Aliases**
Type alias allows you to create a new name for an existing type.

**Example:**

```go
package main

import "fmt"

type Age int  // Type alias for int

func main() {
    var a Age = 30
    fmt.Println(a)  // 30
}
```

#### **b. Type Conversion**
Go allows you to convert between compatible types.

**Example:**

```go
package main

import "fmt"

func main() {
    var x int = 42
    var y float64 = float64(x)  // Convert int to float64
    fmt.Println(y)  // 42.0
}
```

---

### **5. Function Types**

Functions in Go can also be treated as first-class citizens. You can assign a function to a variable, pass functions as arguments, and return functions from other functions.

**Example:**

```go
package main

import "fmt"

// Function that returns another function
func multiply(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

func main() {
    // Create a multiplier function
    double := multiply(2)
    fmt.Println(double(5))  // 10
}
```

---

### **6. Error Type**

Go has a built-in `error` type, which is an interface. It is commonly used for error handling in Go programs.

**Example:**

```go
package main

import "fmt"

func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

func main() {
    result, err := divide(10, 0)
    if err != nil {
        fmt.Println(err)  // cannot divide by zero
    } else {
        fmt.Println(result)
    }
}
```

---

### **Summary**

- **Basic Types**: `int`, `float64`, `bool`, `string`
- **Composite Types**: `array`, `slice`, `map`, `struct`
- **Pointer Types**: `*T`
- **Interface Types**: Used for defining behaviors that types can implement.
- **Type Aliases & Type Conversion**: `type`, explicit type conversion.
- **Function Types**: Functions can be treated as values and passed around.
- **Error Type**: `error` is a built-in interface for handling errors.

Go's type system is designed to be simple but flexible, allowing both high-level abstractions (such as interfaces and structs) and low-level control (such as pointers and type conversions).