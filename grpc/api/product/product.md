Great — let’s build a **product fetching API** in this gRPC + Go + Nuxt setup.

---

## 🎯 Goal

- Add a new gRPC method to fetch products
- Add a REST endpoint in Go to expose it
- Call it from your Nuxt frontend

---

## 🧩 Step 1: Update `.proto` file

Let’s add a `ProductService` to your existing `proto/helloworld.proto`, or better, create a new one:

### 📄 `proto/product.proto`

```proto
syntax = "proto3";

package product;

option go_package = "github.com/yourusername/go-grpc-example/pb";

message Product {
  int32 id = 1;
  string name = 2;
  double price = 3;
}

message ProductListRequest {}
message ProductListResponse {
  repeated Product products = 1;
}

service ProductService {
  rpc ListProducts (ProductListRequest) returns (ProductListResponse);
}
```

---

## 🛠 Step 2: Generate Go code

```bash
protoc \
  --go_out=pb \
  --go-grpc_out=pb \
  --go_opt=paths=source_relative \
  --go-grpc_opt=paths=source_relative \
  proto/product.proto
```

---

## 🧠 Step 3: Implement the Product gRPC Service

### 📁 `server/product_server.go`

```go
package main

import (
    "context"
    pb "github.com/yourusername/go-grpc-example/pb"
)

type productServer struct {
    pb.UnimplementedProductServiceServer
}

func (s *productServer) ListProducts(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductListResponse, error) {
    products := []*pb.Product{
        {Id: 1, Name: "MacBook Pro", Price: 2499.99},
        {Id: 2, Name: "iPhone 15", Price: 1099.99},
        {Id: 3, Name: "AirPods Pro", Price: 249.99},
    }

    return &pb.ProductListResponse{Products: products}, nil
}
```

### 🧩 Update `server/main.go` to register:

```go
pb.RegisterProductServiceServer(grpcServer, &productServer{})
```

---

## 🌍 Step 4: Add REST API endpoint

### 📄 `api/product_api.go`

```go
package main

import (
    "context"
    "encoding/json"
    "net/http"
    "time"

    pb "github.com/yourusername/go-grpc-example/pb"
)

func makeProductHandler(client pb.ProductServiceClient) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        res, err := client.ListProducts(ctx, &pb.ProductListRequest{})
        if err != nil {
            http.Error(w, "Failed to fetch products: "+err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(res.Products)
    }
}
```

### 📄 In `api/main.go`, register this handler:

```go
productClient := pb.NewProductServiceClient(conn)
http.HandleFunc("/api/products", makeProductHandler(productClient))
```

---

## 🧪 Step 5: Test REST API

Start your servers and visit:

```
http://localhost:8080/api/products
```

You’ll get:

```json
[
  { "id": 1, "name": "MacBook Pro", "price": 2499.99 },
  { "id": 2, "name": "iPhone 15", "price": 1099.99 },
  { "id": 3, "name": "AirPods Pro", "price": 249.99 }
]
```

---

## 💻 Step 6: Fetch from Nuxt

### Example: `pages/products.vue`

```vue
<script setup>
const { data, error } = await useFetch('/api/products', {
  baseURL: 'http://localhost:8080',
})
</script>

<template>
  <div>
    <h1>Products</h1>
    <ul v-if="data">
      <li v-for="product in data" :key="product.id">
        {{ product.name }} — ${{ product.price }}
      </li>
    </ul>
    <p v-else-if="error">Failed to load products</p>
    <p v-else>Loading...</p>
  </div>
</template>
```

---

### ✅ You're done!

Now you have:
- A gRPC method to fetch products
- A REST API layer exposing it
- A Nuxt page displaying the data

---

Want to go next-level? I can help with:
- Filtering / search support
- Caching
- Dockerized setup
- Auth (JWT or OAuth)

Let me know what direction you wanna take this 👇