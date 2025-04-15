Ah, gotcha! 👌 Du möchtest also **Swagger/OpenAPI-Dokumentation mit JWT Auth**, **aber ohne Gin**. Kein Problem – das geht natürlich auch!

---

## ✅ Ziel

| Feature               | Umsetzung                              |
|------------------------|------------------------------------------|
| 🔐 Swagger mit JWT Auth | ohne Gin – z. B. mit `net/http` oder anderem Router |
| 📄 OpenAPI Specs       | automatisch generiert mit `swag` CLI |
| 🧪 In Swagger testbar  | „Authorize“-Button mit Bearer Token |

---

## 🔧 Voraussetzungen

- Go REST API (egal ob `net/http`, `chi`, `echo`, `mux`, …)
- Du nutzt `swaggo/swag` zur Doku
- Du willst **JWT** via `Authorization: Bearer <token>` dokumentieren

---

## ✍️ 1. Projekt vorbereiten (falls noch nicht)

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go get github.com/swaggo/http-swagger
go get github.com/swaggo/files
```

Dann `swag init` ausführen.

---

## 📄 2. In `main.go` (Swagger-Header):

```go
// @title        Meine API
// @version      1.0
// @description  REST API mit JWT Auth (ohne Gin)
// @host         localhost:8080
// @BasePath     /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
package main
```

---

## ✍️ 3. Endpunkte dokumentieren

Beispiel:

```go
// @Summary   Alle Produkte abrufen (geschützt)
// @Tags      products
// @Produce   json
// @Success   200 {array} Product
// @Security  BearerAuth
// @Router    /products [get]
func getAllProducts(w http.ResponseWriter, r *http.Request) {
	// Token-Check manuell oder per Middleware
}
```

Die Annotation `@Security BearerAuth` sagt Swagger: dieser Endpoint braucht JWT im Header!

---

## 📚 4. Swagger-UI einbinden (ohne Gin)

Nutze `http.ServeMux` oder `mux.Router`:

```go
import (
    "net/http"
    "github.com/swaggo/http-swagger"
    _ "your/module/docs" // automatisch generiert durch swag init
)

func main() {
    http.HandleFunc("/products", getAllProducts)

    // Swagger UI
    http.Handle("/swagger/", http.StripPrefix("/swagger/", httpSwagger.WrapHandler))

    http.ListenAndServe(":8080", nil)
}
```

Öffne: [http://localhost:8080/swagger/index.html](http://localhost:8080/swagger/index.html)

---

## 🔐 5. In Swagger „Authorize“-Button verwenden

Wenn du alles richtig annotiert hast, zeigt Swagger oben:

🔐 `Authorize`

→ Du gibst dort ein:

```
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6...
```

→ Swagger schickt den Header bei jedem Request mit.

---

## 🧠 Was du bekommst

| Feature                 | Ergebnis                    |
|--------------------------|------------------------------|
| ✅ Auth dokumentiert     | JWT ist sichtbar im UI       |
| ✅ Kein Gin nötig        | Funktioniert mit jedem Router |
| ✅ OpenAPI konform       | Für Postman, Codegen etc.    |

---

## 🔥 Optional

Willst du auch…

- ✅ Swagger Docs als JSON/YAML exportieren?
- ✅ JWT generieren (Login-Endpoint)?
- ✅ Swagger deployen (S3, GitHub Pages etc.)?

Sag Bescheid – wir machen’s rund 😎