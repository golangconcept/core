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
