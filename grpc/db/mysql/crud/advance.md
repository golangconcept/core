Sehr schön! 🧠 Du willst sauberen, **strukturieren Code** mit:

- ✅ **DTO (Data Transfer Object)** – für saubere Kommunikation im gRPC
- ✅ **DAO (Data Access Object)** – Trennung von DB-Logik

Das ist richtig gut für **Clean Architecture**, **Testbarkeit** und **Wartbarkeit** deines gRPC Backends!

---

## 🧱 Aufbau mit DTO & DAO

Hier ist das Konzept:

```
grpc-product-app/
├── proto/                     ← Protobuf-Definitionen
├── dto/                       ← Data Transfer Objects (Request/Response)
│   └── product_dto.go
├── dao/                       ← Data Access Layer (DB-Queries)
│   └── product_dao.go
├── models/                    ← DB-Model (GORM)
│   └── product.go
├── service/
│   └── product_service.go     ← gRPC-Service + Mapping zwischen DTO ↔ DAO ↔ Proto
├── db/
│   └── mysql.go
├── main.go
```

---

## 📄 `models/product.go` – GORM DB-Struktur (Entity)

```go
package models

type Product struct {
    ID       uint   `gorm:"primaryKey"`
    Name     string
    Category string
    Price    float64
}
```

---

## 📄 `dto/product_dto.go` – DTOs für Service Layer

```go
package dto

type ProductDTO struct {
    ID       uint
    Name     string
    Category string
    Price    float64
}

type CreateProductDTO struct {
    Name     string
    Category string
    Price    float64
}

type UpdateProductDTO struct {
    ID       uint
    Name     string
    Category string
    Price    float64
}
```

---

## 📄 `dao/product_dao.go` – DAO für DB-Zugriff

```go
package dao

import (
    "grpc-product-app/dto"
    "grpc-product-app/models"
    "gorm.io/gorm"
)

type ProductDAO struct {
    DB *gorm.DB
}

func (dao *ProductDAO) CreateProduct(input dto.CreateProductDTO) (*models.Product, error) {
    product := models.Product{
        Name:     input.Name,
        Category: input.Category,
        Price:    input.Price,
    }
    if err := dao.DB.Create(&product).Error; err != nil {
        return nil, err
    }
    return &product, nil
}

func (dao *ProductDAO) UpdateProduct(input dto.UpdateProductDTO) (*models.Product, error) {
    var product models.Product
    if err := dao.DB.First(&product, input.ID).Error; err != nil {
        return nil, err
    }

    product.Name = input.Name
    product.Category = input.Category
    product.Price = input.Price

    if err := dao.DB.Save(&product).Error; err != nil {
        return nil, err
    }
    return &product, nil
}

func (dao *ProductDAO) DeleteProduct(id uint) error {
    return dao.DB.Delete(&models.Product{}, id).Error
}

func (dao *ProductDAO) GetByID(id uint) (*models.Product, error) {
    var product models.Product
    if err := dao.DB.First(&product, id).Error; err != nil {
        return nil, err
    }
    return &product, nil
}

func (dao *ProductDAO) GetAll() ([]models.Product, error) {
    var products []models.Product
    err := dao.DB.Find(&products).Error
    return products, err
}

func (dao *ProductDAO) GetByCategory(category string) ([]models.Product, error) {
    var products []models.Product
    err := dao.DB.Where("category = ?", category).Find(&products).Error
    return products, err
}
```

---

## 📄 `service/product_service.go` – gRPC Service mit Mapping

```go
package service

import (
    "context"
    "fmt"
    "grpc-product-app/dao"
    "grpc-product-app/dto"
    "grpc-product-app/proto"
)

type ProductServer struct {
    proto.UnimplementedProductServiceServer
    DAO *dao.ProductDAO
}

func toProto(p *dto.ProductDTO) *proto.Product {
    return &proto.Product{
        Id:       uint32(p.ID),
        Name:     p.Name,
        Category: p.Category,
        Price:    float32(p.Price),
    }
}

func fromModelToDTO(model *models.Product) *dto.ProductDTO {
    return &dto.ProductDTO{
        ID:       model.ID,
        Name:     model.Name,
        Category: model.Category,
        Price:    model.Price,
    }
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *proto.ProductInput) (*proto.ProductResponse, error) {
    input := dto.CreateProductDTO{
        Name:     req.Name,
        Category: req.Category,
        Price:    float64(req.Price),
    }

    product, err := s.DAO.CreateProduct(input)
    if err != nil {
        return nil, err
    }

    return &proto.ProductResponse{Product: toProto(fromModelToDTO(product))}, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, req *proto.ProductUpdate) (*proto.ProductResponse, error) {
    input := dto.UpdateProductDTO{
        ID:       uint(req.Id),
        Name:     req.Name,
        Category: req.Category,
        Price:    float64(req.Price),
    }

    product, err := s.DAO.UpdateProduct(input)
    if err != nil {
        return nil, err
    }

    return &proto.ProductResponse{Product: toProto(fromModelToDTO(product))}, nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, req *proto.ProductIDRequest) (*proto.DeleteResponse, error) {
    err := s.DAO.DeleteProduct(uint(req.Id))
    if err != nil {
        return nil, err
    }

    return &proto.DeleteResponse{
        Message: fmt.Sprintf("Produkt mit ID %d gelöscht", req.Id),
    }, nil
}
```

---

## 🚀 `main.go` – Start mit DAO + Server

```go
func main() {
    db := db.InitMySQL()
    dao := &dao.ProductDAO{DB: db}
    grpcServer := grpc.NewServer()

    proto.RegisterProductServiceServer(grpcServer, &service.ProductServer{
        DAO: dao,
    })

    lis, _ := net.Listen("tcp", ":50051")
    log.Println("gRPC läuft auf Port 50051...")
    grpcServer.Serve(lis)
}
```

---

## ✅ Vorteile von DTO/DAO-Trennung

| Schicht | Aufgabe |
|--------|---------|
| **DTO** | Saubere Datenobjekte für Services/API |
| **DAO** | Zugriff auf Datenbank, gekapselt |
| **Model** | GORM-Struktur für Tabellen |
| **Service** | Business-Logik + gRPC-Handler |

---

## 🔥 Bonus: Willst du auch...

- ✅ REST-Gateway (`grpc-gateway`)
- ✅ Unit Tests mit Mock DAO
- ✅ Swagger-Doku
- ✅ Auth mit gRPC Interceptors

Ich helf dir beim nächsten Step – du bestimmst, wie sauber es werden soll 😎