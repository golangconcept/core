Super, lass uns ganz praktisch durchgehen, **wie du die Middleware mit Rollen-/JWT-Logik korrekt anwendest** 💡  
Wir nehmen an, du hast bereits folgende Bausteine:

- `JWTAuthMiddleware`: prüft Token allgemein ✅  
- `RoleAuthMiddleware("admin")`: prüft Rolle ✅  
- REST-Handler wie `GetAllProducts`, `AdminStats` ✅

---

## 🧭 Ziel

Wir wollen z. B. folgendes Setup in `main.go`:

| Route            | Middleware                         | Zugriff         |
|-------------------|-------------------------------------|------------------|
| `/login`          | – (öffentlich)                     | alle             |
| `/products`       | `JWTAuthMiddleware`                | eingeloggte User |
| `/admin/stats`    | `RoleAuthMiddleware("admin")`      | nur Admins       |

---

## 🧱 So baust du’s ein: `main.go`

```go
package main

import (
	"net/http"
	"grpc-product-app/auth"
	"grpc-product-app/rest"
)

func main() {
	// Öffentliche Route
	http.HandleFunc("/login", rest.LoginHandler)

	// Geschützte Route für eingeloggte Nutzer (z. B. user oder admin)
	http.Handle("/products", auth.JWTAuthMiddleware(http.HandlerFunc(rest.GetAllProducts)))

	// Nur Admins dürfen zugreifen
	http.Handle("/admin/stats", auth.RoleAuthMiddleware("admin")(http.HandlerFunc(rest.AdminStats)))

	// Server starten
	http.ListenAndServe(":8080", nil)
}
```

---

## ✅ Beispiel für Handler (geschützt)

```go
// rest/products.go
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(auth.UserContextKey()).(string)
	w.Write([]byte("Hallo, " + username + "! Hier sind deine Produkte."))
}
```

---

## 🔐 Für `/admin/stats`:

```go
func AdminStats(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(auth.UserContextKey()).(string)
	w.Write([]byte("Adminbereich für " + username))
}
```

---

## 🧪 Test mit `curl`

### 1. Login und Token holen

```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"username":"admin","password":"adminpass"}'
```

→ Response:

```json
{"token": "eyJhbGciOiJIUzI1..."}
```

---

### 2. Zugriff mit Token

```bash
curl -H "Authorization: Bearer eyJhbGciOi..." http://localhost:8080/products
```

oder:

```bash
curl -H "Authorization: Bearer eyJhbGciOi..." http://localhost:8080/admin/stats
```

---

## 💬 Was passiert bei:

| Fall                          | Ergebnis                     |
|-------------------------------|------------------------------|
| Kein Token                    | 401 Unauthorized             |
| Token, aber falsche Rolle     | 403 Forbidden                |
| Gültiger Token & Rolle        | 200 OK + Inhalt              |

---

## 🔧 Optional: Middleware debuggen

Willst du Debug-Output wie `fmt.Println("Token geprüft: ...")` einbauen?  
Kein Problem, einfach in der Middleware hinzufügen zum besseren Logging.

---

## 🔥 Nächste Ausbaustufen?

- ✅ Zugriff auf `userID`, `email` aus JWT?
- ✅ Scopes wie `["read:orders", "edit:users"]`?
- ✅ Routen dynamisch schützen (z. B. per Config)?

Sag Bescheid – ich helf dir beim nächsten Step! 😎