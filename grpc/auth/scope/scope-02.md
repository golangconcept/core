Perfekt, jetzt machen wir dein **Auth-System richtig mächtig**:  
👉 **Rollen + Scopes kombiniert**, also z. B.:

> Ein Admin darf *alles* (durch Rolle), ein Editor darf nur bestimmte Scopes wie `"edit:products"`, und normale User z. B. `"read:orders"`.

---

## 🧩 Ziel-Struktur

Dein JWT sieht dann z. B. so aus:

```json
{
  "username": "lena",
  "role": "editor",
  "scopes": ["edit:products", "read:products"],
  "exp": 1713000000
}
```

---

## ✅ Was wir jetzt bauen:

| Feature                  | Funktion                            |
|--------------------------|--------------------------------------|
| 🔐 Rolle prüfen            | z. B. `"admin"`                     |
| 🎯 Scopes prüfen           | z. B. `"read:orders"`              |
| ➕ Oder beides kombinieren | Admin darf alles, sonst Scope nötig |

---

## 📄 JWT-Erzeugung anpassen

### `auth/jwt.go`

```go
func GenerateJWT(username, role string, scopes []string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"role":     role,
		"scopes":   scopes,
		"exp":      time.Now().Add(time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
```

---

## 🔐 Kombinierte Middleware: `RoleOrScopeAuthMiddleware`

### 📄 `auth/role_scope_middleware.go`

```go
func RoleOrScopeAuthMiddleware(allowedRoles []string, requiredScope string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Context().Value(jwtTokenContextKey).(*jwt.Token)

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// ✅ 1. Check Role
			if role, ok := claims["role"].(string); ok {
				for _, allowed := range allowedRoles {
					if role == allowed {
						next.ServeHTTP(w, r)
						return
					}
				}
			}

			// ✅ 2. Check Scope (wenn Rolle nicht passt)
			if scopes, ok := claims["scopes"].([]interface{}); ok {
				for _, s := range scopes {
					if scopeStr, ok := s.(string); ok && scopeStr == requiredScope {
						next.ServeHTTP(w, r)
						return
					}
				}
			}

			// ❌ Keine Rolle + kein Scope → verbieten
			http.Error(w, "Forbidden – insufficient permissions", http.StatusForbidden)
		}))
	}
}
```

---

## 🧩 Middleware anwenden

```go
// Admins oder User mit Scope "edit:products"
http.Handle("/products/edit",
	auth.RoleOrScopeAuthMiddleware([]string{"admin"}, "edit:products")(
		http.HandlerFunc(rest.EditProduct),
	),
)

// Admins oder Scopes mit "read:orders"
http.Handle("/orders",
	auth.RoleOrScopeAuthMiddleware([]string{"admin"}, "read:orders")(
		http.HandlerFunc(rest.GetOrders),
	),
)
```

---

## 🧪 Beispiel JWTs

1. ✅ **Admin**
```json
{ "role": "admin", "scopes": [] }
```

2. ✅ **Editor mit Scope**
```json
{ "role": "editor", "scopes": ["edit:products"] }
```

3. ❌ **User ohne Berechtigung**
```json
{ "role": "user", "scopes": ["read:profile"] }
```

---

## 🧠 Warum kombinieren?

| Vorteil                         | Beschreibung                            |
|----------------------------------|------------------------------------------|
| ✅ **Rollen für grobe Kontrolle** | Admins/Moderatoren                      |
| ✅ **Scopes für Details**         | CRUD-Operationen gezielt erlauben       |
| ✅ **Erweiterbar**                | Superadmin, API-Nutzer, Drittanbieter…  |

---

## 🧪 Willst du Unit Tests dafür?

Klar! Du kannst mit `httptest.NewRequest()` + `Recorder` einfach testen:
- Token mit admin-Rolle → ✅ erlaubt
- Token mit passendem Scope → ✅ erlaubt
- Kein Scope, keine Rolle → ❌ verboten

Ich kann dir die Testdatei sofort mitschreiben!

---

## 📄 Bonus: Swagger mit Rollen & Scopes

Mit OpenAPI kannst du sogar definieren:
- Welche Scopes pro Endpoint notwendig sind
- „Authorize“-Button mit Scopes

Wenn du magst, zeig ich dir das auch 🚀

---

## ✍️ Nächster Schritt?

Sag mir, ob du willst:

- [ ] 🧪 Unit Tests
- [ ] 🛡️ Scope-Gruppen (z. B. `"products:*"`)
- [ ] 🔄 Refresh Token
- [ ] 🧠 Benutzerrechte aus DB laden statt im Token

Ich helf dir direkt weiter!