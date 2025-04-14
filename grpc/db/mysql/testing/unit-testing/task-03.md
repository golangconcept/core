Yes, let’s go full **REST CRUD** mit 🔥 **POST, PUT und DELETE** + Unit Tests mit Mock DAO!  
Wir machen das clean, testbar und stabil – ganz ohne Datenbank 💪

---

## 🧱 Projektstruktur-Erweiterung

```
rest/
├── product_handler.go         // ✅ REST-API Handler
├── product_handler_test.go    // ✅ Unit Tests mit Mock DAO
```

---

## 🔁 1. Neue REST-Endpunkte (POST, PUT, DELETE)

### 📄 `rest/product_handler.go`

Erweitere deinen `ProductHandler`:

```go
func (h *ProductHandler) RegisterRoutes(r *gin.Engine) {
    r.GET("/products", h.GetAll)
    r.GET("/products/:id", h.GetByID)
    r.POST("/products", h.Create)
    r.PUT("/products/:id", h.Update)
    r.DELETE("/products/:id", h.Delete)
}

func (h *ProductHandler) Create(c *gin.Context) {
    var input struct {
        Name     string  `json:"name"`
        Category string  `json:"category"`
        Price    float64 `json:"price"`
    }
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product, err := h.DAO.CreateProduct(dto.CreateProductDTO{
        Name: input.Name, Category: input.Category, Price: input.Price,
    })
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) Update(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    var input struct {
        Name     string  `json:"name"`
        Category string  `json:"category"`
        Price    float64 `json:"price"`
    }
    if err := c.BindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    product, err := h.DAO.UpdateProduct(dto.UpdateProductDTO{
        ID:       uint(id),
        Name:     input.Name,
        Category: input.Category,
        Price:    input.Price,
    })

    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) Delete(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
        return
    }

    err = h.DAO.DeleteProduct(uint(id))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Deleted"})
}
```

---

## 🧪 2. Unit Tests für POST, PUT, DELETE

### 📄 `rest/product_handler_test.go` – Tests erweitern

Import & Setup bleibt gleich – du brauchst nur neue Tests 👇

---

### ✅ POST /products

```go
func TestCreateProduct(t *testing.T) {
    r := setupRouter()

    body := `{"name":"Regal","category":"Möbel","price":59.99}`
    req, _ := http.NewRequest("POST", "/products", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, 201, w.Code)
    assert.Contains(t, w.Body.String(), "Regal")
}
```

---

### ✅ PUT /products/:id

```go
func TestUpdateProduct_Success(t *testing.T) {
    r := setupRouter()

    body := `{"name":"Neuer Stuhl","category":"Möbel","price":66.6}`
    req, _ := http.NewRequest("PUT", "/products/1", strings.NewReader(body))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "Neuer Stuhl")
}
```

---

### ✅ DELETE /products/:id

```go
func TestDeleteProduct_Success(t *testing.T) {
    r := setupRouter()

    req, _ := http.NewRequest("DELETE", "/products/1", nil)
    w := httptest.NewRecorder()
    r.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
    assert.Contains(t, w.Body.String(), "Deleted")
}
```

---

## 🧠 Warum das wichtig ist

| Vorteil           | Bedeutung                                  |
|-------------------|---------------------------------------------|
| **POST/PUT/DELETE getestet** | Keine bösen Überraschungen im Live-System ✅ |
| **Mocks statt DB**          | Schneller, isolierter Testlauf ⏱️ |
| **Fehlerfälle simulierbar** | Z. B. `404`, `400`, `500` ⚠️ |

---

## 🔥 Optional: Noch mehr Features?

- ✅ **Swagger/OpenAPI** Doku generieren?
- ✅ **Integration Tests** mit echtem MySQL in Docker?
- ✅ **GitHub Actions** für automatische Tests?

Sag einfach Bescheid – wir bauen das aus wie du willst 😎