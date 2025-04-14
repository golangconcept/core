Sehr nice! 🧪 Zeit für **Unit Tests mit einem Mock-DAO** – damit du deine gRPC-Service-Logik **ohne echte Datenbank** testen kannst. Genau das ist der Vorteil von DTO/DAO-Trennung 💡

---

## ✅ Ziel

- Teste deinen `ProductServer` (gRPC Service) isoliert
- Simuliere DAO-Funktionen mit einem **Mock**
- Kein DB-Zugriff, kein GORM nötig

---

## 🧱 Verzeichnisstruktur (Tests)

```
grpc-product-app/
├── service/
│   ├── product_service.go
│   └── product_service_test.go  ✅
```

---

## 📦 1. Testing + Mock Packages installieren

```bash
go get github.com/stretchr/testify
```

---

## 👨‍🔧 2. Interface für DAO (für Mocks)

Erstelle oder erweitere `dao/product_dao.go`:

```go
package dao

import (
    "grpc-product-app/dto"
    "grpc-product-app/models"
)

type ProductDAOInterface interface {
    CreateProduct(input dto.CreateProductDTO) (*models.Product, error)
    UpdateProduct(input dto.UpdateProductDTO) (*models.Product, error)
    DeleteProduct(id uint) error
    GetByID(id uint) (*models.Product, error)
    GetAll() ([]models.Product, error)
    GetByCategory(category string) ([]models.Product, error)
}
```

In deinem Service dann statt `DAO *ProductDAO`, nun:

```go
type ProductServer struct {
    proto.UnimplementedProductServiceServer
    DAO dao.ProductDAOInterface
}
```

---

## 🧪 3. Unit Test mit Fake/Mock DAO (`product_service_test.go`)

```go
package service_test

import (
    "context"
    "errors"
    "testing"

    "github.com/stretchr/testify/assert"
    "grpc-product-app/dto"
    "grpc-product-app/models"
    "grpc-product-app/proto"
    "grpc-product-app/service"
)

// 👉 Fake DAO für Tests
type mockProductDAO struct{}

func (m *mockProductDAO) CreateProduct(input dto.CreateProductDTO) (*models.Product, error) {
    return &models.Product{ID: 1, Name: input.Name, Category: input.Category, Price: input.Price}, nil
}
func (m *mockProductDAO) UpdateProduct(input dto.UpdateProductDTO) (*models.Product, error) {
    if input.ID == 999 {
        return nil, errors.New("not found")
    }
    return &models.Product{ID: input.ID, Name: input.Name, Category: input.Category, Price: input.Price}, nil
}
func (m *mockProductDAO) DeleteProduct(id uint) error {
    if id == 999 {
        return errors.New("not found")
    }
    return nil
}
func (m *mockProductDAO) GetByID(id uint) (*models.Product, error) {
    if id == 42 {
        return &models.Product{ID: 42, Name: "Stuhl", Category: "Möbel", Price: 29.99}, nil
    }
    return nil, errors.New("not found")
}
func (m *mockProductDAO) GetAll() ([]models.Product, error) {
    return []models.Product{
        {ID: 1, Name: "Tisch", Category: "Möbel", Price: 99.99},
        {ID: 2, Name: "Lampe", Category: "Licht", Price: 19.99},
    }, nil
}
func (m *mockProductDAO) GetByCategory(category string) ([]models.Product, error) {
    if category == "Möbel" {
        return []models.Product{
            {ID: 1, Name: "Tisch", Category: "Möbel", Price: 99.99},
        }, nil
    }
    return []models.Product{}, nil
}
```

---

## 🧪 4. Beispiel-Tests

```go
func TestGetProductByID_Success(t *testing.T) {
    server := service.ProductServer{DAO: &mockProductDAO{}}
    res, err := server.GetProductByID(context.Background(), &proto.ProductIDRequest{Id: 42})

    assert.NoError(t, err)
    assert.Equal(t, "Stuhl", res.Product.Name)
    assert.Equal(t, float32(29.99), res.Product.Price)
}

func TestGetProductByID_NotFound(t *testing.T) {
    server := service.ProductServer{DAO: &mockProductDAO{}}
    _, err := server.GetProductByID(context.Background(), &proto.ProductIDRequest{Id: 123})

    assert.Error(t, err)
}

func TestCreateProduct(t *testing.T) {
    server := service.ProductServer{DAO: &mockProductDAO{}}
    req := &proto.ProductInput{
        Name:     "Regal",
        Category: "Möbel",
        Price:    59.99,
    }

    res, err := server.CreateProduct(context.Background(), req)
    assert.NoError(t, err)
    assert.Equal(t, "Regal", res.Product.Name)
    assert.Equal(t, float32(59.99), res.Product.Price)
}
```

---

## ✅ TL;DR

| Technik         | Beschreibung                              |
|------------------|-------------------------------------------|
| `interface` für DAO | Ermöglicht flexible Implementierung     |
| Fake DAO         | Liefert kontrollierte Rückgaben           |
| `stretchr/testify` | Für elegante Assertions & Fehlerprüfungen |

---

## 🔥 Bonus

Willst du:

- deinen DAO mit **Mockgen** automatisch mocken?
- auch **REST Layer** mit Mocks testen?
- Integration Tests mit echtem MySQL in Docker?

Sag Bescheid – ich bin ready 😎