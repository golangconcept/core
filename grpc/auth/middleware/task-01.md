Perfekt! 🔐 Lass uns eine **vollständige, robuste JWT Middleware** in Go bauen – ganz ohne Framework wie Gin. Die Middleware schützt deine REST-Endpunkte und prüft automatisch:

- Ist ein **Authorization Header** da?
- Ist der **JWT gültig & nicht abgelaufen**?
- Optional: Extrahiere **Claims** wie `username` aus dem Token

---

## 🧱 Voraussetzungen

- Standard `net/http`
- JWT mit `github.com/golang-jwt/jwt/v5` installiert:
  
```bash
go get github.com/golang-jwt/jwt/v5
```

---

## 🔧 Schritt 1: JWT Secret & Generieren (Kurzfassung)

In `auth/jwt.go`:

```go
package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("mein_super_geheimes_token") // in .env oder config auslagern

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour).Unix(), // 1 Stunde gültig
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
```

---

## 🔐 Schritt 2: Die Middleware

In `auth/middleware.go`:

```go
package auth

import (
	"context"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const userContextKey = contextKey("username")

func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Header lesen
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing or invalid Authorization header", http.StatusUnauthorized)
			return
		}

		// Token extrahieren
		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		// Token validieren
		token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Optional: Claims auslesen
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			username := claims["username"].(string)

			// In Request-Context legen (für spätere Handler)
			ctx := context.WithValue(r.Context(), userContextKey, username)
			r = r.WithContext(ctx)
		}

		// Weiterreichen an echten Handler
		next.ServeHTTP(w, r)
	})
}
```

---

## 📥 Zugriff auf User im Handler

In deinem Handler (z. B. `/products`), der durch die Middleware geht:

```go
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value(auth.UserContextKey())
	w.Write([]byte("Hallo, " + username.(string)))
}
```

Und `auth/jwt.go` bekommt noch eine Getter-Funktion:

```go
func UserContextKey() interface{} {
	return userContextKey
}
```

---

## 🔧 Schritt 3: Middleware registrieren

In deinem `main.go`:

```go
import (
	"net/http"
	"grpc-product-app/auth"
	"grpc-product-app/rest"
)

func main() {
	http.HandleFunc("/login", rest.LoginHandler)

	protected := http.NewServeMux()
	protected.HandleFunc("/products", rest.GetAllProducts)

	http.Handle("/products", auth.JWTAuthMiddleware(protected))

	http.ListenAndServe(":8080", nil)
}
```

---

## ✅ So funktioniert’s jetzt:

- `/login` → gibt JWT zurück
- `/products` → geht nur mit gültigem `Authorization: Bearer <token>`
- User-Info steckt im `context`

---

## 🧪 Unit Tests für die Middleware?

Haben wir schon gemacht (siehe vorheriger Schritt), aber sag Bescheid, wenn du auch z. B. abgelaufene Tokens oder manipulierte Claims testen willst.

---

## 🔥 Optional ausbauen?

- ✅ Claims wie `roles`, `userID`, `email` hinzufügen?
- ✅ Token Refresh Endpoint?
- ✅ Token-Blacklist (Logout)?
- ✅ Integration mit Swagger UI?

Sag’s einfach – ich helf dir beim nächsten Schritt 😎