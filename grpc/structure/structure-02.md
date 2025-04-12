Ah, perfect — if your folder is named `server/` instead of `internal/`, no problem at all! I’ll adjust everything to fit that layout while **preserving clean structure** and **SOLID principles**.

---

## 🗂️ Updated Folder Structure (with `server/` as root)

```
go-grpc-example/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point (clean main)
├── server/
│   ├── cache/
│   │   └── redis_cache.go          # Redis logic
│   ├── db/
│   │   └── mongo_product_repo.go   # MongoDB logic
│   ├── grpc/
│   │   └── product_server.go       # gRPC handler
│   ├── config/
│   │   ├── redis.go                # Redis setup
│   │   ├── mongo.go                # Mongo setup
│   │   └── grpc.go                 # gRPC server setup
│   └── listener/
│       └── serve.go                # Start gRPC server
├── pb/
│   └── (generated .pb.go files)
├── proto/
│   └── product.proto               # Protobuf definitions
├── go.mod
├── go.sum
└── docker-compose.yml
```

---

## 🔧 Adjusted `cmd/server/main.go`

```go
package main

import (
	"context"
	"go-grpc-example/server/config"
	"go-grpc-example/server/listener"
)

func main() {
	ctx := context.Background()

	redisClient := config.SetupRedis()
	mongoColl := config.SetupMongo(ctx)
	grpcServer := config.SetupGRPCServer(redisClient, mongoColl)

	listener.Start(grpcServer)
}
```

---

### 🔁 Folder/Package Changes

Everything that was in `internal/` just moves to `server/`:

| Old Path                          | New Path                         |
|----------------------------------|----------------------------------|
| `internal/cache/redis_cache.go`  | `server/cache/redis_cache.go`    |
| `internal/db/mongo_repo.go`      | `server/db/mongo_repo.go`        |
| `internal/grpc/product_server.go`| `server/grpc/product_server.go`  |
| `internal/server/grpc.go`        | `server/config/grpc.go`          |
| `internal/server/mongo.go`       | `server/config/mongo.go`         |
| `internal/server/redis.go`       | `server/config/redis.go`         |
| `internal/server/listener.go`    | `server/listener/serve.go`       |

> ✅ The package names inside files stay consistent with their folders, like `package cache`, `package db`, etc.

---

### ✅ Imports Update Example

In `grpc.go` (inside `server/config`):

```go
import (
	"go-grpc-example/server/cache"
	"go-grpc-example/server/db"
	"go-grpc-example/server/grpc"
)
```

In `main.go`:

```go
import (
	"go-grpc-example/server/config"
	"go-grpc-example/server/listener"
)
```

---

## 💡 Why This Still Follows SOLID & Clean Code

- **Single Responsibility**: each file/folder does one thing
- **Dependency Inversion**: cache and DB access via interfaces
- **Extensibility**: want to swap Mongo with Postgres? Just implement a new `ProductRepository`
- **Testability**: mock interfaces, test logic in isolation
- **Clarity**: easy to onboard new devs and scale

---

## ✅ You’re Ready to Ship

- Everything neat under `server/`
- Easily expandable (REST, Nuxt, tests, workers, etc.)
- Fully Docker-compatible

---

If you want, I can:

- Generate the full working project
- Add REST API gateway for Nuxt
- Add seeding support for Mongo
- Create a `Makefile` or scripts to simplify running

Just say the word — or tell me what you want next 😎