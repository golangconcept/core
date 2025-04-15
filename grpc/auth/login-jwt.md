Jawoll! 🔐 Lass uns einen **Login-Endpoint bauen**, der bei erfolgreichem Login ein **JWT (JSON Web Token)** zurückgibt.

Das ist super für Auth in REST-APIs – und perfekt kombinierbar mit deinem Swagger & Unit Testing Setup!

---

## ✅ Ziel

| Feature             | Funktion                                  |
|---------------------|--------------------------------------------|
| 🔐 Login-Endpoint    | POST `/login` mit Benutzername/Passwort   |
| 🧾 JWT zurückgeben   | Als Bearer Token (im Header verwendbar)   |
| ✅ Sicher signiert    | Mit Secret und Ablaufzeit (exp)           |

---

## 📦 Benötigte Libs

Installiere JWT-Bibliothek:

```bash
go get github.com/golang-jwt/jwt/v5
```

---

## 📄 1. JWT Utility

### `auth/jwt.go`

```go
package auth

import (
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte("mein_super_geheimes_token") // 🔐 Sicher aufbewahren!

func GenerateJWT(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 1).Unix(), // 1h gültig
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
```

---

## 📄 2. Login-Handler

### `rest/auth_handler.go`

```go
package rest

import (
	"encoding/json"
	"net/http"
	"grpc-product-app/auth"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var creds LoginRequest
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil || creds.Username == "" || creds.Password == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// ⚠️ Beispiel: Fake User Check – hier z. B. DB-Query oder Hardcode
	if creds.Username != "admin" || creds.Password != "pass123" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	token, err := auth.GenerateJWT(creds.Username)
	if err != nil {
		http.Error(w, "Error creating token", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"token": token})
}
```

---

## 📄 3. Route registrieren

### In `main.go` oder Router-Setup:

```go
http.HandleFunc("/login", rest.LoginHandler)
```

---

## 🔍 Test mit `curl` oder Swagger

```bash
curl -X POST http://localhost:8080/login \
     -H "Content-Type: application/json" \
     -d '{"username":"admin","password":"pass123"}'
```

💥 Response:

```json
{"token":"eyJhbGciOiJIUzI1NiIsInR5cCI6..."}
```

Jetzt kannst du diesen Token als `Authorization: Bearer <token>` an deine geschützten Endpunkte schicken.

---

## 🛡️ Optional: Middleware zum Prüfen

JWT Middleware in z. B. `auth/middleware.go`:

```go
func JWTAuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

		token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// ✅ Token gültig – weiter geht’s
		next(w, r)
	}
}
```

Dann bei deiner geschützten Route:

```go
http.HandleFunc("/products", auth.JWTAuthMiddleware(rest.GetAll))
```

---

## 🔐 JWT im Swagger dokumentieren?

Wenn du willst, bau ich das gleich passend in deine Swagger-Doku ein → `Authorize` Button etc.  
Einfach sagen 😉

---

## 🧪 Willst du...

- ✅ Unit-Tests für Login + Middleware?
- ✅ Passwort-Hashing mit `bcrypt`?
- ✅ Token Refresh & Logout?

Ich helf dir gerne dabei – sag, was du brauchst 😎