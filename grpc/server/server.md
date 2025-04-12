Awesome — setting up a **Go project with gRPC** is a solid move if you're building scalable microservices. Here's a complete setup from scratch:

---

## ⚙️ Go gRPC Project Setup Guide

### 🧱 Tech Stack
- Go (≥1.18)
- gRPC (via Protocol Buffers)
- `protoc` compiler
- `protoc-gen-go` and `protoc-gen-go-grpc` plugins
- Optional: `Makefile` or shell scripts for easier builds

---

### 📁 Project Structure

```
go-grpc-example/
│
├── proto/
│   └── helloworld.proto
│
├── server/
│   └── main.go
│
├── client/
│   └── main.go
│
├── pb/                     # Auto-generated Go files
│
├── go.mod
└── README.md
```

---

### 1. 🧰 Install Required Tools

#### Install Protocol Buffer Compiler

- macOS: `brew install protobuf`
- Ubuntu: `sudo apt install -y protobuf-compiler`
- Windows: [Download here](https://github.com/protocolbuffers/protobuf/releases)

#### Install Go plugins

```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

Make sure `$GOPATH/bin` is in your `$PATH`.

---

### 2. 📦 Initialize the Go module

```bash
mkdir go-grpc-example && cd go-grpc-example
go mod init github.com/yourusername/go-grpc-example
```

---

### 3. 📜 Create a .proto file

**`proto/helloworld.proto`**

```proto
syntax = "proto3";

package helloworld;

option go_package = "github.com/yourusername/go-grpc-example/pb";

service Greeter {
  rpc SayHello (HelloRequest) returns (HelloReply);
}

message HelloRequest {
  string name = 1;
}

message HelloReply {
  string message = 1;
}
```

---

### 4. ⚙️ Generate Go code

```bash
protoc \
  --go_out=pb \
  --go-grpc_out=pb \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  proto/helloworld.proto
```

This generates two files:
- `pb/helloworld.pb.go`
- `pb/helloworld_grpc.pb.go`

---

### 5. 🚀 Implement Server

**`server/main.go`**

```go
package main

import (
    "context"
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/yourusername/go-grpc-example/pb"
)

type greeterServer struct {
    pb.UnimplementedGreeterServer
}

func (s *greeterServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
    return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterGreeterServer(grpcServer, &greeterServer{})

    log.Println("Server listening at :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

---

### 6. 🧪 Implement Client

**`client/main.go`**

```go
package main

import (
    "context"
    "log"
    "time"

    "google.golang.org/grpc"
    pb "github.com/yourusername/go-grpc-example/pb"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewGreeterClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Go Developer"})
    if err != nil {
        log.Fatalf("could not greet: %v", err)
    }

    log.Printf("Greeting: %s", resp.Message)
}
```

---

### ✅ Run It!

```bash
# In one terminal
go run server/main.go

# In another terminal
go run client/main.go
```

---

### 🛠 Bonus: Add Makefile

```makefile
PROTOC=protoc
PROTO_SRC=proto
PROTO_DST=pb

proto:
	$(PROTOC) \
	--go_out=$(PROTO_DST) \
	--go-grpc_out=$(PROTO_DST) \
	--go_opt=paths=source_relative \
	--go-grpc_opt=paths=source_relative \
	$(PROTO_SRC)/helloworld.proto
```

Run `make proto` to regenerate code after changes.

---

Would you like me to zip this up for you as a ready-to-run template or help you add features like streaming or TLS?