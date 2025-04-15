Perfekt! 🔥 Lass uns jetzt **Unit Tests** für deinen **JWT-Login Endpoint** und die **Auth Middleware** bauen. So stellen wir sicher, dass nur gültige Logins durchkommen – und dein System 🔐 sicher & stabil läuft.

---

## ✅ Was wir testen:

| Testbereich         | Beschreibung                            |
|----------------------|------------------------------------------|
| 🔐 Login-Endpoint     | Korrekte/fehlerhafte Logins             |
| 🛡️ JWT Middleware     | Token vorhanden? gültig? abgelaufen?    |

---

## 📁 Projektstruktur (vereinfacht)

```
auth/
├── jwt.go              // Token erzeugen & prüfen
├── middleware.go       // HTTP Middleware
├── jwt_test.go         // ✅ Unit Tests für JWT/Middleware

rest/
├── auth_handler.go     // /login Endpoint
├── auth_handler_test.go // ✅ Unit Tests für Login
```

---

## 🔧 Voraussetzungen

```bash
go get github.com/stretchr/testify
```

---

## 🧪 1. `auth_handler_test.go`: Login testen

```go
package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoginHandler_Success(t *testing.T) {
	payload := map[string]string{
		"username": "admin",
		"password": "pass123",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	LoginHandler(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Contains(t, rec.Body.String(), "token")
}

func TestLoginHandler_Unauthorized(t *testing.T) {
	payload := map[string]string{
		"username": "admin",
		"password": "wrongpass",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	LoginHandler(rec, req)

	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
```

✅ Du testest erfolgreichen & fehlerhaften Login isoliert – ohne DB.

---

## 🧪 2. `jwt_test.go`: Token erzeugen und Middleware testen

### 📄 `auth/jwt_test.go`

```go
package auth

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	token, err := GenerateJWT("tester")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)
}
```

---

### 📄 `auth/middleware_test.go`

```go
func TestJWTAuthMiddleware_ValidToken(t *testing.T) {
	token, _ := GenerateJWT("tester")

	handler := JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	handler(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)
}

func TestJWTAuthMiddleware_MissingToken(t *testing.T) {
	handler := JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()

	handler(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}

func TestJWTAuthMiddleware_InvalidToken(t *testing.T) {
	handler := JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer abc123.invalid.token")
	rec := httptest.NewRecorder()

	handler(rec, req)
	assert.Equal(t, http.StatusUnauthorized, rec.Code)
}
```

---

## ✅ Tests laufen lassen

```bash
go test ./... -v
```

---

## 🧠 Bonus-Tipp: Middleware über mehrere Tests wiederverwenden

```go
func makeRequestWithToken(t *testing.T, token string) *httptest.ResponseRecorder {
	handler := JWTAuthMiddleware(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	rec := httptest.NewRecorder()

	handler(rec, req)
	return rec
}
```

Dann:

```go
rec := makeRequestWithToken(t, "dein-token")
assert.Equal(t, 401, rec.Code)
```

---

## 🔥 Optional ausbauen?

- ✅ Token-Claims im Kontext speichern (`r.Context()`)
- ✅ Ablaufzeit (`exp`) testen (z. B. manipuliert)
- ✅ Integrationstests über HTTP-Client + Login + Zugriff

Sag Bescheid – ich helfe dir, was du brauchst! 😎