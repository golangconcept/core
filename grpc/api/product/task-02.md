Perfekt! Du willst eine **gRPC API in Go**, die Daten aus **MySQL** abruft – und zwar für:

✅ **Produkte**  
✅ **Kategorien**

Ich zeig dir gleich den **kompletten Ablauf**:
- gRPC mit `.proto` Dateien für `ProductService` und `CategoryService`  
- MySQL-Datenbank  
- gRPC-Server in Go  
- Daten aus MySQL lesen  
- Docker-ready Struktur auf Wunsch

---

## 📦 Beispiel-Datenmodell (MySQL)

### Tabelle: `categories`

```sql
CREATE TABLE categories (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255)
);
```

### Tabelle: `products`

```sql
CREATE TABLE products (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(255),
  price DECIMAL(10,2),
  category_id INT,
  FOREIGN KEY (category_id) REFERENCES categories(id)
);
```

---

## 📄 1. `proto/product.proto`

```proto
syntax = "proto3";

package pb;

service ProductService {
  rpc ListProducts (Empty) returns (ProductListResponse);
}

message Product {
  int32 id = 1;
  string name = 2;
  float price = 3;
  int32 category_id = 4;
}

message ProductListResponse {
  repeated Product products = 1;
}

message Empty {}
```

---

## 📄 2. `proto/category.proto`

```proto
syntax = "proto3";

package pb;

service CategoryService {
  rpc ListCategories (Empty) returns (CategoryListResponse);
}

message Category {
  int32 id = 1;
  string name = 2;
}

message CategoryListResponse {
  repeated Category categories = 1;
}
```

Dann:  
```bash
protoc --go_out=. --go-grpc_out=. proto/*.proto
```

---

## 🧠 3. MySQL-Verbindung (Go)

### 📁 Datei: `db/db.go`

```go
package db

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func NewMySQL() (*sql.DB, error) {
    dsn := "user:password@tcp(mysql:3306)/yourdb"
    return sql.Open("mysql", dsn)
}
```

---

## 🧩 4. Produkt-Service mit MySQL

### 📁 Datei: `internal/product/service.go`

```go
package product

import (
    "context"
    "database/sql"
    "yourmodule/pb"
)

type Service struct {
    DB *sql.DB
}

func (s *Service) ListProducts(ctx context.Context, _ *pb.Empty) (*pb.ProductListResponse, error) {
    rows, err := s.DB.Query("SELECT id, name, price, category_id FROM products")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var products []*pb.Product
    for rows.Next() {
        var p pb.Product
        if err := rows.Scan(&p.Id, &p.Name, &p.Price, &p.CategoryId); err != nil {
            continue
        }
        products = append(products, &p)
    }

    return &pb.ProductListResponse{Products: products}, nil
}
```

> Das gleiche Prinzip gilt für `CategoryService`.

---

## 🏁 5. `main.go` – gRPC Server starten

```go
package main

import (
    "log"
    "net"

    "google.golang.org/grpc"
    "yourmodule/db"
    pb "yourmodule/pb"
    "yourmodule/internal/product"
)

func main() {
    listener, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    mysqlDB, err := db.NewMySQL()
    if err != nil {
        log.Fatalf("failed to connect to DB: %v", err)
    }

    grpcServer := grpc.NewServer()
    productService := &product.Service{DB: mysqlDB}
    pb.RegisterProductServiceServer(grpcServer, productService)

    log.Println("✅ gRPC server listening on :50051")
    grpcServer.Serve(listener)
}
```

---

## 🐳 6. Docker Setup (Optional)

### `docker-compose.yml`

```yaml
version: "3.9"
services:
  mysql:
    image: mysql:8
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: yourdb
      MYSQL_USER: user
      MYSQL_PASSWORD: password
    ports:
      - "3306:3306"

  product-service:
    build: ./product-service
    ports:
      - "50051:50051"
    depends_on:
      - mysql
```

---

## ✅ Was du bekommst

| Feature             | Status       |
|---------------------|--------------|
| gRPC Server         | ✅ Ready      |
| Protobuf Services   | ✅ Category & Product |
| MySQL Integration   | ✅ Ja         |
| Docker Support      | ✅ Optional   |
| Erweiterbar         | ✅ Sicher     |

---

## 👉 Nächster Schritt?

Ich kann dir jetzt:

- 🔧 Ein fertiges Template auf GitHub oder ZIP machen  
- 🌐 Einen REST-Gateway oder GraphQL oben drauf bauen  
- 🧪 Ein Test-Schema + Seed-Daten dazupacken  
- 📦 Makefile und proto-builder dazulegen

Sag einfach:  
**"Mach ZIP fertig"**  
oder  
**"Mach GitHub Template mit beiden Services"**  

Ready wenn du's bist! 💪