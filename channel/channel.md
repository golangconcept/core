# channel in golang

In Go, channels are a powerful feature of the language's concurrency model, enabling goroutines to `communicate with each other and synchronize execution`.

- Channels provide a way for one goroutine to send data to another, allowing them to coordinate without the need for explicit locks or shared memory.

They are a core part of Go's approach to concurrency, which is based on the `CSP` (**Communicating Sequential Processes**) model.

### Key Concepts About Channels

- `Channels` are used to send and receive data between goroutines.
- They are typed, meaning you define a channel with a specific data type (chan int, chan string, etc.).
- **Channels are blocking by default, meaning that**:
  - If a goroutine tries to send data on a channel and no other goroutine is ready to receive it, it will block until another goroutine receives the data.
  - If a goroutine tries to receive data from a channel and no data is available, it will block until data is sent.

## Basic Syntax of Channels

1. Creating a channel:

```go
// Create an unbuffered channel for int data
ch := make(chan int)
```

2. Sending data to a channel:

```go
ch <- 5 // Send value 5 to the channel
```

3. Receiving data from a channel:

```go
x := <-ch // Receive value from the channel and assign it to x
```

4. Closing a channel:

```go
close(ch) // Close the channel
```

When a channel is closed, `any subsequent sends will panic`. However, reads from a closed channel will continue to work until the channel is empty, and the ok value will be false when no more values are available.

```go
v, ok := <-ch // ok is false if the channel is closed and empty
```
