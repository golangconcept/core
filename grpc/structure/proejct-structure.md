Perfect! Let’s now **organize your Go project** into a clean folder structure with properly separated files and packages – ideal for maintainability and following the **SOLID** principles. 🚀

---

## 🗂️ Recommended Folder Structure

```
go-grpc-example/
├── cmd/
│   └── server/
│       └── main.go                 # Entry point (clean main)
├── internal/
│   ├── cache/
│   │   └── redis_cache.go          # Redis cache logic
│   ├── db/
│   │   └── mongo_product_repo.go   # MongoDB product repo
│   ├── grpc/
│   │   └── product_server.go       # gRPC server implementation
│   └── server/
│       ├── grpc.go                 # gRPC server setup
│       ├── redis.go                # Redis setup
│       ├── mongo.go                # Mongo setup
│       └── listener.go             # Listener + Serve
├── pb/
│   └── (generated .pb.go files)
├── proto/
│   └── product.proto               # Protobuf definition
├── go.mod
├── go.sum
└── docker-compose.yml
```

---

## 🔧 Step-by-Step Breakdown

### 1. `cmd/server/main.go`

```go
package main

import (
	"context"
	"log"

	"go-grpc-example/internal/server"
)

func main() {
	ctx := context.Background()

	redisClient := server.SetupRedis()
	mongoColl := server.SetupMongo(ctx)
	grpcServer := server.SetupGRPCServer(redisClient, mongoColl)

	server.StartListener(grpcServer)
}
```

---

### 2. `internal/server/redis.go`

```go
package server

import "github.com/redis/go-redis/v9"

func SetupRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "redis:6379",
	})
}
```

---

### 3. `internal/server/mongo.go`

```go
package server

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongo(ctx context.Context) *mongo.Collection {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://mongo:27017"))
	if err != nil {
		log.Fatalf("MongoDB connection failed: %v", err)
	}
	return client.Database("mydb").Collection("products")
}
```

---

### 4. `internal/server/grpc.go`

```go
package server

import (
	"time"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
	"go-grpc-example/internal/cache"
	"go-grpc-example/internal/db"
	"go-grpc-example/internal/grpc"
	"go-grpc-example/pb"
	"google.golang.org/grpc"
)

func SetupGRPCServer(redisClient *redis.Client, mongoColl *mongo.Collection) *grpc.Server {
	repo := &db.MongoProductRepo{Coll: mongoColl}
	cache := &cache.RedisProductCache{Client: redisClient, TTL: 5 * time.Minute}
	productService := &grpc.ProductServer{Repo: repo, Cache: cache}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, productService)

	return grpcServer
}
```

---

### 5. `internal/server/listener.go`

```go
package server

import (
	"log"
	"net"

	"google.golang.org/grpc"
)

func StartListener(grpcServer *grpc.Server) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	log.Println("🚀 gRPC server listening on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
```

---

### 6. `internal/db/mongo_product_repo.go`

```go
package db

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go-grpc-example/pb"
)

type MongoProductRepo struct {
	Coll *mongo.Collection
}

func (m *MongoProductRepo) GetAll(ctx context.Context) ([]*pb.Product, error) {
	cursor, err := m.Coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var products []*pb.Product
	for cursor.Next(ctx) {
		var doc struct {
			ID    int32   `bson:"id"`
			Name  string  `bson:"name"`
			Price float64 `bson:"price"`
		}
		if err := cursor.Decode(&doc); err == nil {
			products = append(products, &pb.Product{
				Id:    doc.ID,
				Name:  doc.Name,
				Price: doc.Price,
			})
		}
	}

	return products, nil
}
```

---

### 7. `internal/cache/redis_cache.go`

```go
package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	"go-grpc-example/pb"
)

type RedisProductCache struct {
	Client *redis.Client
	TTL    time.Duration
}

func (r *RedisProductCache) Get(ctx context.Context, key string) (*pb.ProductListResponse, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}

	var res pb.ProductListResponse
	if err := json.Unmarshal([]byte(val), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (r *RedisProductCache) Set(ctx context.Context, key string, data *pb.ProductListResponse) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return r.Client.Set(ctx, key, bytes, r.TTL).Err()
}
```

---

### 8. `internal/grpc/product_server.go`

```go
package grpc

import (
	"context"
	"log"

	"go-grpc-example/pb"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]*pb.Product, error)
}

type ProductCache interface {
	Get(ctx context.Context, key string) (*pb.ProductListResponse, error)
	Set(ctx context.Context, key string, data *pb.ProductListResponse) error
}

type ProductServer struct {
	pb.UnimplementedProductServiceServer
	Repo  ProductRepository
	Cache ProductCache
}

func (s *ProductServer) ListProducts(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductListResponse, error) {
	const key = "products"

	if cached, err := s.Cache.Get(ctx, key); err == nil {
		log.Println("✅ Redis cache hit")
		return cached, nil
	}

	log.Println("❌ Cache miss — fetching from MongoDB")
	products, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := &pb.ProductListResponse{Products: products}
	_ = s.Cache.Set(ctx, key, res)
	return res, nil
}
```

---

## ✅ What You’ve Achieved

✅ **Clean folder structure**  
✅ **SOLID-based logic separation**  
✅ **Reusable interfaces for DB/cache**  
✅ **gRPC logic isolated from startup**  
✅ **Great testability & maintainability**

---

## 📦 Want a ready-to-go zip or GitHub repo?

Let me know and I’ll generate the full working example with Docker, proto files, and everything pre-wired — so you can just run `docker-compose up` and go 🚀

Want Nuxt frontend or REST layer added too? Let's do it!