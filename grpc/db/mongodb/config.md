Sweet! Let’s pull product data **directly from MongoDB** instead of hardcoding it — and still support caching with Redis for speed. 🔥

---

## 🎯 What You’ll Do

1. Connect Go to MongoDB (using `mongo-go-driver`)
2. Fetch products from a `products` collection
3. Cache the result in Redis
4. Return it via gRPC and REST as before

---

## 📦 Tech Stack

- MongoDB: holds product data
- Redis: caches product data
- gRPC + Go: backend logic
- REST: API gateway

---

## 🧱 Step 1: Add MongoDB to Docker Compose

Update your `docker-compose.yml`:

```yaml
services:
  mongo:
    image: mongo
    ports:
      - "27017:27017"
    volumes:
      - mongo-data:/data/db
    container_name: mongo-db

volumes:
  mongo-data:
```

---

## 📦 Step 2: Install MongoDB Go Driver

```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

---

## 🔧 Step 3: Connect to MongoDB in `ProductServer`

### Modify `ProductServer` struct

```go
type ProductServer struct {
    pb.UnimplementedProductServiceServer
    redis     *redis.Client
    mongoColl *mongo.Collection
}
```

---

### Update `NewProductServer()`

```go
func NewProductServer(redisAddr, mongoURI string) *ProductServer {
    ctx := context.Background()

    // Redis client
    rdb := redis.NewClient(&redis.Options{
        Addr: redisAddr,
    })

    // MongoDB client
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
    if err != nil {
        log.Fatalf("Mongo connection error: %v", err)
    }

    coll := client.Database("mydb").Collection("products")

    return &ProductServer{
        redis:     rdb,
        mongoColl: coll,
    }
}
```

---

## 📦 Step 4: Fetch from MongoDB in `ListProducts`

```go
func (s *ProductServer) ListProducts(ctx context.Context, req *pb.ProductListRequest) (*pb.ProductListResponse, error) {
    const key = "products"

    // Check Redis cache
    if val, err := s.redis.Get(ctx, key).Result(); err == nil {
        var cached pb.ProductListResponse
        if err := json.Unmarshal([]byte(val), &cached); err == nil {
            log.Println("✅ Redis cache hit")
            return &cached, nil
        }
    }

    log.Println("❌ Cache miss — fetching from MongoDB")

    // Fetch from MongoDB
    cursor, err := s.mongoColl.Find(ctx, bson.M{})
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

    res := &pb.ProductListResponse{Products: products}

    // Cache in Redis
    if jsonData, err := json.Marshal(res); err == nil {
        s.redis.Set(ctx, key, jsonData, 5*time.Minute)
    }

    return res, nil
}
```

---

## 🛠 Update `main.go`

```go
pb.RegisterProductServiceServer(grpcServer, server.NewProductServer("redis:6379", "mongodb://mongo:27017"))
```

✅ This connects Go to both Redis and Mongo inside Docker.

---

## 📦 Step 5: Insert Sample Products in MongoDB

You can do this via a script, Mongo shell, or UI like MongoDB Compass.

Example:

```json
[
  { "id": 1, "name": "MacBook Pro", "price": 2499.99 },
  { "id": 2, "name": "iPhone 15", "price": 1099.99 },
  { "id": 3, "name": "AirPods Pro", "price": 249.99 }
]
```

Run:

```bash
docker exec -it mongo-db mongosh
use mydb
db.products.insertMany([
  { id: 1, name: "MacBook Pro", price: 2499.99 },
  { id: 2, name: "iPhone 15", price: 1099.99 },
  { id: 3, name: "AirPods Pro", price: 249.99 }
])
```

---

## ✅ Test It

1. Start services:

```bash
docker-compose up --build
```

2. Call your REST endpoint:

```
http://localhost:8080/api/products
```

First call → MongoDB fetch  
Next calls → Redis cache 🚀

---

## 🎁 Bonus Ideas

- Add MongoDB indexes for performance
- Support filters (e.g., category, price range)
- Use Redis for per-user caching
- Add admin API to refresh cache

---

Want help writing a seeder or integrating this with Nuxt for real-time product lists? I got you!