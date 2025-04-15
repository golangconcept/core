Jawoll! 🔐 Lass uns den **Auth Header mit JWT** sauber in Swagger dokumentieren – damit Swagger-UI, Frontend-Teams, Postman und Co. genau wissen, wie man deine API **authentifiziert**.

---

## ✅ Ziel

| Feature                  | Was passiert                                  |
|--------------------------|-----------------------------------------------|
| 🔐 JWT Header             | Swagger erwartet `Authorization: Bearer <jwt>` |
| 📄 Dokumentiert in UI    | Sichtbar in Swagger unter „Authorize“ Button |
| 🧪 Testbar               | Du kannst direkt in Swagger JWT einfügen     |

---

## 🧱 Voraussetzungen

- Du nutzt Gin + Swagger (`swaggo`)
- Dein API verwendet JWT mit `Authorization: Bearer <token>`
- Du willst bestimmte Endpunkte absichern

---

## ✍️ 1. Swagger Security Definition hinzufügen

In deiner `main.go` Datei, direkt oben:

```go
// @title        Produkt-API
// @version      1.0
// @description  API mit JWT Authentifizierung
// @host         localhost:8080
// @BasePath     /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
```

---

## 🔐 2. Endpunkte mit Auth annotieren

In deinem Handler-File `rest/product_handler.go` (für geschützte Endpunkte):

```go
// @Summary Geschützte Produktliste
// @Tags products
// @Produce json
// @Success 200 {array} models.Product
// @Security BearerAuth
// @Router /products [get]
func (h *ProductHandler) GetAll(c *gin.Context) { ... }
```

Du kannst `@Security BearerAuth` bei allen Endpunkten hinzufügen, die JWT brauchen.

---

## 💡 Hinweis zu Middleware

Damit es funktioniert, musst du in deiner Gin-API natürlich **eine JWT-Middleware haben**, z. B.:

```go
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" || !strings.HasPrefix(token, "Bearer ") {
            c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            return
        }

        // Token prüfen (z.B. mit jwt-go) ...

        c.Next()
    }
}
```

Dann schützt du deine Routen:

```go
auth := r.Group("/")
auth.Use(AuthMiddleware())
auth.GET("/products", h.GetAll)
```

---

## 🧪 3. Swagger ausprobieren

1. Swagger starten: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)  
2. Klicke oben rechts auf 🔒 **Authorize**
3. Gib ein:
   ```
   Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6...
   ```
4. Jetzt kannst du **geschützte Endpunkte** direkt testen 🎯

---

## ✅ Vorteile

| Vorteil                | Warum es geil ist                     |
|------------------------|----------------------------------------|
| 🔐 Doku + Test in einem | Keine Doku mehr schreiben nötig        |
| 🚀 Klar für Frontend    | Jeder weiß: Bearer Token nötig         |
| 🤖 Kompatibel mit Postman | Swagger JSON enthält alles für Tests |

---

## 🔥 Optional

Willst du auch:

- ✅ JWT generieren und refreshen?
- ✅ Swagger global hosten (z. B. über S3)?
- ✅ Swagger für deine gRPC-API mit grpc-gateway?

Sag’s einfach – ich baue es mit dir durch 😎