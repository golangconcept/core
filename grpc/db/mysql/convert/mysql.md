Perfekt! Lass uns ein **vollständiges Beispiel** bauen – von Repository über Service bis zum gRPC-Handler.  
Wir setzen das Ganze mit **MySQL**, **Redis-Caching**, **gRPC**, **SOLID-Struktur** und einem sauberen Aufruffluss um.

---

## 🧱 Struktur (Mini-Microservice: `product-service`)

```
product-service/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── db/                  # Repository layer
│   │   ├── interface.go
│   │   └── mysql_product_repo.go
│   ├── cache/
│   │   └── redis.go
│   ├── service/             # Business logic
│   │   └── product_service.go
│   └── handler/             # gRPC handler
│       └── product_handler.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── db/
│   │   └── mysql.go
│   └── redis/
│       └── redis.go
├── proto/
│   └── product.proto
├── go.mod
└── docker-compose.yml
```

---

## 🛠️ 1. `proto/product.proto`

```proto
syntax = "proto3";

package pb;

service ProductService {
  rpc ListProducts (Empty) returns (ProductListResponse);
}

message Empty {}

message Product {
  int32 id = 1;
  string name = 2;
  float price = 3;
}

message ProductListResponse {
  repeated Product products = 1;
}
```

Dann generieren:

```bash
protoc --go_out=. --go-grpc_out=. proto/product.proto
```

---

## 🧩 2. Repository Layer (`internal/db/mysql_product_repo.go`)

```go
package db

import (
	"context"
	"database/sql"
	pb "product-service/proto"
)

type ProductRepository interface {
	GetAll(ctx context.Context) ([]*pb.Product, error)
}

type MySQLProductRepo struct {
	DB *sql.DB
}

func (m *MySQLProductRepo) GetAll(ctx context.Context) ([]*pb.Product, error) {
	rows, err := m.DB.QueryContext(ctx, `SELECT id, name, price FROM products`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []*pb.Product
	for rows.Next() {
		var p pb.Product
		if err := rows.Scan(&p.Id, &p.Name, &p.Price); err == nil {
			products = append(products, &p)
		}
	}
	return products, rows.Err()
}
```

---

## 🚀 3. Redis Cache (`internal/cache/redis.go`)

```go
package cache

import (
	"context"
	"encoding/json"
	"time"

	"github.com/redis/go-redis/v9"
	pb "product-service/proto"
)

type ProductCache struct {
	Client *redis.Client
	TTL    time.Duration
}

func (c *ProductCache) Get(ctx context.Context, key string) (*pb.ProductListResponse, error) {
	val, err := c.Client.Get(ctx, key).Result()
	if err != nil {
		return nil, err
	}
	var res pb.ProductListResponse
	if err := json.Unmarshal([]byte(val), &res); err != nil {
		return nil, err
	}
	return &res, nil
}

func (c *ProductCache) Set(ctx context.Context, key string, data *pb.ProductListResponse) error {
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.Client.Set(ctx, key, bytes, c.TTL).Err()
}
```

---

## 🧠 4. Business Layer (`internal/service/product_service.go`)

```go
package service

import (
	"context"
	"log"

	"product-service/internal/cache"
	"product-service/internal/db"
	pb "product-service/proto"
)

type ProductService struct {
	Repo  db.ProductRepository
	Cache *cache.ProductCache
}

func (s *ProductService) ListProducts(ctx context.Context) (*pb.ProductListResponse, error) {
	const key = "products"

	// Redis Hit
	if cached, err := s.Cache.Get(ctx, key); err == nil {
		log.Println("✅ Redis HIT")
		return cached, nil
	}

	// Redis Miss
	log.Println("❌ Cache MISS, DB-Fetch")
	products, err := s.Repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.ProductListResponse{Products: products}

	_ = s.Cache.Set(ctx, key, resp)
	return resp, nil
}
```

---

## 🎯 5. gRPC Handler (`internal/handler/product_handler.go`)

```go
package handler

import (
	"context"

	"product-service/internal/service"
	pb "product-service/proto"
)

type ProductHandler struct {
	Service *service.ProductService
	pb.UnimplementedProductServiceServer
}

func (h *ProductHandler) ListProducts(ctx context.Context, req *pb.Empty) (*pb.ProductListResponse, error) {
	return h.Service.ListProducts(ctx)
}
```

---

## 🏁 6. Main (`cmd/server/main.go`)

```go
package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"product-service/internal/cache"
	"product-service/internal/db"
	"product-service/internal/handler"
	"product-service/internal/service"
	"product-service/pkg/config"
	dbInit "product-service/pkg/db"
	redisInit "product-service/pkg/redis"
	pb "product-service/proto"
)

func main() {
	cfg := config.Load()

	// DB init
	dbConn, err := dbInit.NewMySQL(cfg.MySQL)
	if err != nil {
		log.Fatalf("MySQL error: %v", err)
	}
	defer dbConn.Close()

	// Redis init
	redisClient, err := redisInit.NewRedis(cfg.Redis)
	if err != nil {
		log.Fatalf("Redis error: %v", err)
	}

	// Compose all
	repo := &db.MySQLProductRepo{DB: dbConn}
	cache := &cache.ProductCache{Client: redisClient, TTL: cfg.CacheTTL}
	svc := &service.ProductService{Repo: repo, Cache: cache}
	handler := &handler.ProductHandler{Service: svc}

	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, handler)

	log.Println("✅ gRPC Server läuft auf :50051")
	grpcServer.Serve(lis)
}
```

---

## 🧪 7. `.env` (optional)

```
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=secret
MYSQL_DB=shopdb

REDIS_ADDR=localhost:6379
CACHE_TTL=300
```

---

## ✅ Ergebnis:

🔁 `GET /ListProducts` (via gRPC)  
➡️ checkt Redis  
➡️ fallback auf MySQL  
➡️ Antwort + Cache-Speicherung  

---

## 📦 Nächstes?

Möchtest du:

- Eine ZIP-Datei mit allem Code?  
- REST-Gateway hinzufügen?  
- Eine Kategorie-Entity genauso aufbauen?  
- Tests mit `sqlmock` oder Redis-Mock?

Sag einfach z. B. **"Mach mir auch Kategorie-Service"** oder **"REST-Gateway für das Ganze"**  
Dann baue ich das sofort mit rein! 🚀