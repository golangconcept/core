Klar! 😎 Ich erweitere dein gRPC-Backend um die klassischen CRUD-Funktionen:

✅ **Produkt einfügen**  
✅ **Produkt bearbeiten (Update)**  
✅ **Produkt löschen (Delete)**  

Alle mit GORM & MySQL – direkt im bestehenden gRPC-Setup.

---

## 📦 1. ✏️ `product.proto` um neue RPCs erweitern

```proto
// proto/product.proto

service ProductService {
  rpc ListProducts (Empty) returns (ProductList) {}
  rpc GetProductByID (ProductIDRequest) returns (ProductResponse) {}
  rpc GetProductsByCategory (CategoryRequest) returns (ProductList) {}

  // Neue Methoden 👇
  rpc CreateProduct (ProductInput) returns (ProductResponse) {}
  rpc UpdateProduct (ProductUpdate) returns (ProductResponse) {}
  rpc DeleteProduct (ProductIDRequest) returns (DeleteResponse) {}
}

message ProductInput {
  string name = 1;
  string category = 2;
  float price = 3;
}

message ProductUpdate {
  uint32 id = 1;
  string name = 2;
  string category = 3;
  float price = 4;
}

message DeleteResponse {
  string message = 1;
}
```

---

## 🔁 2. Protobuf neu generieren

```bash
protoc --go_out=. --go-grpc_out=. proto/product.proto
```

---

## 🛠 3. Neue Methoden in `product_service.go` implementieren

```go
func (s *ProductServer) CreateProduct(ctx context.Context, req *proto.ProductInput) (*proto.ProductResponse, error) {
    product := models.Product{
        Name:     req.Name,
        Category: req.Category,
        Price:    float64(req.Price),
    }
    if err := s.DB.Create(&product).Error; err != nil {
        return nil, err
    }

    return &proto.ProductResponse{
        Product: &proto.Product{
            Id:       uint32(product.ID),
            Name:     product.Name,
            Category: product.Category,
            Price:    float32(product.Price),
        },
    }, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, req *proto.ProductUpdate) (*proto.ProductResponse, error) {
    var product models.Product
    if err := s.DB.First(&product, req.Id).Error; err != nil {
        return nil, err
    }

    product.Name = req.Name
    product.Category = req.Category
    product.Price = float64(req.Price)

    if err := s.DB.Save(&product).Error; err != nil {
        return nil, err
    }

    return &proto.ProductResponse{
        Product: &proto.Product{
            Id:       uint32(product.ID),
            Name:     product.Name,
            Category: product.Category,
            Price:    float32(product.Price),
        },
    }, nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, req *proto.ProductIDRequest) (*proto.DeleteResponse, error) {
    if err := s.DB.Delete(&models.Product{}, req.Id).Error; err != nil {
        return nil, err
    }

    return &proto.DeleteResponse{
        Message: fmt.Sprintf("Produkt mit ID %d gelöscht", req.Id),
    }, nil
}
```

---

## 🧪 4. Beispiele zum Testen (via `grpcurl`)

```bash
# Produkt einfügen
grpcurl -plaintext -d '{"name":"Stuhl", "category":"Möbel", "price":49.99}' localhost:50051 product.ProductService/CreateProduct

# Produkt aktualisieren
grpcurl -plaintext -d '{"id":1, "name":"Neuer Stuhl", "category":"Möbel", "price":59.99}' localhost:50051 product.ProductService/UpdateProduct

# Produkt löschen
grpcurl -plaintext -d '{"id":1}' localhost:50051 product.ProductService/DeleteProduct
```

---

## ✅ TL;DR

| Funktion      | RPC                             | DB-Methode (GORM)    |
|---------------|----------------------------------|-----------------------|
| Einfügen      | `CreateProduct(ProductInput)`    | `db.Create(...)`      |
| Bearbeiten    | `UpdateProduct(ProductUpdate)`   | `db.Save(...)`        |
| Löschen       | `DeleteProduct(ProductIDRequest)`| `db.Delete(...)`      |

---

Wenn du willst, zeig ich dir als Nächstes:

- ✅ Input-Validierung
- ✅ REST-Gateway (`grpc-gateway`)
- ✅ Authentifizierung mit Tokens
- ✅ Logging / Middleware

Sag einfach, wie weit du's treiben willst 😎