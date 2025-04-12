Alles klar! Du möchtest `productServer` und seine `ListProducts`-Methode **in eine separate Datei** auslagern — sauber getrennt, top Idee 💡

Ich zeig dir, wie du das ganz sauber aufteilst:

---

## ✅ Ziel: `product_server.go` auslagern

**Projektstruktur:**
```
go-grpc-example/
├── server/
│   ├── main.go                 # Startet den gRPC-Server
│   └── product_server.go       # Enthält productServer-Implementierung
├── proto/
│   └── product.proto
├── pb/
│   └── product.pb.go
│   └── product_grpc.pb.go
```

---

### 🧠 1. `server/product_server.go`

```go
package server

import (
    "context"

    pb "github.com/yourusername/go-grpc-example/pb"
)

type ProductServer struct {
    pb.UnimplementedProductServiceServer
}

func (s *ProductServer) ListProducts(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductListResponse, error) {
    return &pb.ProductListResponse{
        Products: []*pb.Product{
            {Id: 1, Name: "MacBook Pro", Price: 2499.99},
            {Id: 2, Name: "iPhone 15", Price: 1099.99},
        },
    }, nil
}
```

---

### 🚀 2. `server/main.go`

```go
package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    pb "github.com/yourusername/go-grpc-example/pb"
    "github.com/yourusername/go-grpc-example/server"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    pb.RegisterProductServiceServer(grpcServer, &server.ProductServer{}) // <- hier!

    log.Println("gRPC server listening on :50051")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
```

> 🔁 Vergiss nicht, deine `go.mod` und Imports korrekt zu halten – z. B. `module github.com/yourusername/go-grpc-example`

---

## 📦 Bonus-Tipp: Package `server` verwenden

Du kannst alles in `package server` packen (z. B. für `product_server.go`) und dann einfach importieren in `main.go`.

---

## ✅ Ergebnis

- Code ist sauber getrennt
- `main.go` bleibt schlank
- Du kannst später weitere Services wie `UserServer` etc. hinzufügen

---

Wenn du magst, kann ich dir auch ein fertiges Template-Projekt auf GitHub hochladen oder zum Download als ZIP vorbereiten. Willst du das?