Nice! So you're building a **Go gRPC backend** and want a **Nuxt frontend** to fetch data from it. Since browsers can’t directly talk to gRPC (it’s not HTTP/REST), we’ll add a **REST API layer in Go** to act as a bridge.

Let me break this down for you.

---

## 🌉 Architecture Overview

```
[Nuxt Frontend] ⇄ HTTP (REST) ⇄ [Go REST API Gateway] ⇄ gRPC ⇄ [Go gRPC Server]
```

---

## ⚙️ Step-by-Step: Add a REST API to Fetch Data for Nuxt

---

### 🔁 Step 1: Add REST API in Go (using `net/http`)

We’ll write a simple HTTP API endpoint in Go that calls the gRPC client internally.

#### 🗂 Structure Addition

```
go-grpc-example/
├── api/             # NEW
│   └── main.go      # REST API server
```

---

### 🔧 `api/main.go`

```go
package main

import (
    "context"
    "encoding/json"
    "log"
    "net/http"
    "time"

    "google.golang.org/grpc"
    pb "github.com/yourusername/go-grpc-example/pb"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect to gRPC server: %v", err)
    }
    defer conn.Close()

    client := pb.NewGreeterClient(conn)

    http.HandleFunc("/api/hello", func(w http.ResponseWriter, r *http.Request) {
        name := r.URL.Query().Get("name")
        if name == "" {
            http.Error(w, "Missing 'name' query param", http.StatusBadRequest)
            return
        }

        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
        if err != nil {
            http.Error(w, "gRPC error: "+err.Error(), http.StatusInternalServerError)
            return
        }

        json.NewEncoder(w).Encode(map[string]string{
            "message": resp.Message,
        })
    })

    log.Println("REST API listening on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
```

---

### 🧪 Test the REST API

Run it:

```bash
go run api/main.go
```

Test in browser or Postman:

```
http://localhost:8080/api/hello?name=Nuxt
```

You’ll get:

```json
{ "message": "Hello Nuxt" }
```

---

## 🌐 Step 2: Fetch from Nuxt (Frontend)

Let’s say you’re using **Nuxt 3** with Composition API.

### Example API call:

#### `pages/index.vue`

```vue
<script setup>
const { data, error } = await useFetch('/api/hello', {
  baseURL: 'http://localhost:8080', // or use runtime config
  params: { name: 'Nuxt User' },
})
</script>

<template>
  <div>
    <h1>gRPC API says:</h1>
    <p v-if="data">{{ data.message }}</p>
    <p v-else-if="error">Error: {{ error.message }}</p>
    <p v-else>Loading...</p>
  </div>
</template>
```

---

## ✅ Summary

- Nuxt calls `/api/hello?name=...` via REST
- Go REST API bridges that to the gRPC backend
- gRPC server handles the actual business logic

---

### 🧱 Want extras?
I can help you:
- Dockerize this setup
- Use **gRPC-Gateway** (auto-generates REST for gRPC)
- Add TLS for security
- Add more advanced gRPC features (streaming, auth, etc.)

Want to go further?