Oh yes! 🔥 **REST Layer mit Mocks testen** ist super wichtig – vor allem wenn du z. B. eine **Gin**-API oder ähnliches verwendest.

Wir machen jetzt:

✅ REST-Endpunkte (Gin)  
✅ Unit-Tests mit **Mock DAO**  
✅ Ohne echte Datenbank – nur isoliert logik testen

---

## 🧱 Aufbau (REST Layer über gRPC/DAO/Service)

```
grpc-product-app/
├── rest/
│   ├── product_handler.go      ✅ REST-Endpunkte
│   └── product_handler_test.go ✅ Tests mit Mocks
```

---

## 📄 Beispiel: `rest/product_handler.go`

Hier binden wir unseren DAO/Service im REST-Kontext ein – z. B. mit **Gin**:

```go
package rest

import (
    "github.com/gin-gonic/gin"
    "grpc-product-app/dao"
    "net/http"
    "strconv"
)

type ProductHandler struct {
    DAO dao.ProductDAOInterface
}

func (h *ProductHandler) RegisterRoutes(r *gin.Engine) {
    r.GET("/products", h.GetAll)
    r.GET("/products/:id", h.GetByID)
}

func (h *ProductHandler) GetAll(c *gin.Context) {
    products, err := h.DAO.GetAll()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) GetByID(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    product, err := h.DAO.GetByID(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}
```

---

## 🧪 `rest/product_handler_test.go`

### 🔧 Vorbereitungen

```bash
go get github.com/stretchr/testify
go get github.com/gin-gonic/gin
go get net/http/httptest
```

### ✏️ Test mit Mock DAO + HTTP Requests

```go
package rest_test

import (
    "errors"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "grpc-product-app/models"
    "grpc-product-app/rest"
    "net/http"
    "net/http/httptest"
    "testing"
)

// 🚀 Fake DAO
type mockDAO struct{}

func (m *mockDAO) GetAll() ([]models.Product, error) {
    return []models.Product{{ID: 1, Name: "Stuhl", Category: "Möbel", Price: 49.99}}, nil
}
func (m *mockDAO) GetByID(id uint) (*models.Product, error) {
    if id == 1 {
        return &models.Product{ID: 1, Name: "Stuhl", Category: "Möbel", Price: 49.99}, nil
    }
    return nil, errors.New("not found")
}
// Dummy-Funktionen für nicht genutzte DAO-Methoden:
func (m *mockDAO) CreateProduct(i interface{}) (*models.Product, error) { return nil, nil }
func (m *mockDAO) UpdateProduct(i interface{}) (*models.Product, error) { return nil, nil }
func (m *mockDAO) DeleteProduct(id uint) error                          { return nil }
func (m *mockDAO) GetByCategory(cat string) ([]models.Product, error)   { return nil, nil }

func setupRouter() *gin.Engine {
    gin.SetMode(gin.TestMode)
    router := gin.Default()

    handler := &rest.ProductHandler{DAO: &mockDAO{}}
    handler.RegisterRoutes(router)

    return router
}

func TestGetAllProducts(t *testing.T) {
    r := setupRouter()

    req, _ := http.NewRequest("GET", "/products", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "Stuhl")
}

func TestGetByID_Success(t *testing.T) {
    r := setupRouter()

    req, _ := http.NewRequest("GET", "/products/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "Stuhl")
}

func TestGetByID_NotFound(t *testing.T) {
    r := setupRouter()

    req, _ := http.NewRequest("GET", "/products/999", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, 404, w.Code)
}
```

---

## ✅ Testziel erreicht

| Testfall                  | Ergebnis   |
|---------------------------|------------|
| `/products`               | ✅ 200 + JSON |
| `/products/1`             | ✅ 200 mit Produkt |
| `/products/999`           | ✅ 404 Not Found |
| DAO: keine DB nötig       | ✅ mockDAO simuliert alles |

---

## 🧠 Warum das super ist:

| Vorteil               | Erklärung |
|------------------------|-----------|
| Schnelle Tests         | Kein DB-Setup nötig |
| Fehlerhandling testbar | Simuliere DAO-Fehler |
| REST isoliert          | REST unabhängig von DB/GORM |
| CI/CD ready            | Stabil & reproduzierbar |

---

## 🔥 Nächster Step?

Willst du auch...

- ✅ POST/PUT/DELETE im REST mit Tests?
- ✅ OpenAPI-Schema mit Swagger?
- ✅ Automatische Tests per GitHub Actions?

Sag einfach Bescheid – ich helf dir auf Wunsch weiter 🚀