Top! 🚀 Jetzt machen wir's richtig granular mit **Scopes** wie `"read:orders"` oder `"edit:users"`.

> **Scopes** erlauben feiner abgestimmte **Berechtigungen** als nur `"role": "admin"` – z. B. pro API-Operation. Perfekt für APIs mit unterschiedlichen Nutzerrechten.

---

## 🧠 Was sind Scopes?

**Scopes** sind einfach Strings im JWT, die erlaubte Aktionen beschreiben, z. B.:

```json
{
  "username": "anna",
  "scopes": ["read:products", "delete:orders"]
}
```

So kannst du genau sagen:  
🔹 *Darf sie Produkte lesen?*  
🔸 *Darf sie Bestellungen löschen?*

---

## 🧱 JWT erweitern mit Scopes

### 📄 `auth/jwt.go`

```go
func GenerateJWT(username string, scopes []string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"scopes":   scopes, // 👈 hier!
		"exp":      time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}
```

---

## 🔐 Neue Middleware: `ScopeAuthMiddleware`

### 📄 `auth/scope_middleware.go`

```go
func ScopeAuthMiddleware(requiredScope string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return JWTAuthMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Context().Value(jwtTokenContextKey).(*jwt.Token)

			claims, ok := token.Claims.(jwt.MapClaims)
			if !ok || !token.Valid {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			scopes, ok := claims["scopes"].([]interface{})
			if !ok {
				http.Error(w, "No scopes found", http.StatusForbidden)
				return
			}

			// prüfen ob requiredScope vorhanden ist
			for _, scope := range scopes {
				if scopeStr, ok := scope.(string); ok && scopeStr == requiredScope {
					next.ServeHTTP(w, r)
					return
				}
			}

			http.Error(w, "Forbidden – missing scope", http.StatusForbidden)
		}))
	}
}
```

---

## 🧩 Anwenden in `main.go`

```go
http.Handle("/orders", auth.ScopeAuthMiddleware("read:orders")(http.HandlerFunc(rest.GetOrders)))
http.Handle("/users/edit", auth.ScopeAuthMiddleware("edit:users")(http.HandlerFunc(rest.EditUser)))
```

---

## 🧪 Testen mit Token

Stelle sicher, dass dein JWT beim Login erzeugt wird z. B.:

```go
GenerateJWT("anna", []string{"read:orders", "edit:users"})
```

Dann:

```bash
curl -H "Authorization: Bearer <dein_token>" http://localhost:8080/orders
```

---

## 💡 Vorteile von Scopes

| Vorteil                 | Beschreibung                            |
|--------------------------|------------------------------------------|
| ✅ Feingranular           | „Darf Produkt löschen“ ≠ „lesen“        |
| ✅ Gut für Microservices | Jeder Service prüft nur seine Scopes    |
| ✅ Kombinierbar mit Rollen | Rolle `"admin"` + Scopes wie `"*"`     |

---

## 🧠 Erweiterungsideen

- ✅ Mehrere Scopes auf einmal erlauben (`requiredScopes []string`)
- ✅ Role + Scopes kombinieren
- ✅ Scopes im Swagger sichtbar machen

---

## 💬 Willst du jetzt...

- 🔐 Login erweitern für Scope-Auswahl?
- 🧪 Unit Tests für `ScopeAuthMiddleware`?
- 📄 Swagger/OpenAPI mit Scope-Beschreibung?

Sag einfach Bescheid – ich helf dir Schritt für Schritt 💪