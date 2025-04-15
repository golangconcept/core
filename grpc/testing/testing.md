Klaro! 🚀 Lass uns über **Unit Testing in Go** reden – eines der stärksten Features von Go für sauberen, robusten Code 💪

---

## ✅ Was ist Unit Testing?

Ein **Unit Test** testet **eine Funktion oder Methode isoliert**, ohne Abhängigkeiten wie Datenbank, Netzwerk, etc.

Du testest also nur **die Logik** – ob sie tut, was sie soll.

---

## 📦 Go Built-in Testing Tool

Go bringt alles schon mit – kein Framework nötig!  
Du brauchst nur:

```go
import "testing"
```

Tests laufen mit:

```bash
go test ./...
```

---

## 🧪 Beispiel 1: Einfache Funktion testen

### 📄 `math.go`

```go
package math

func Add(a, b int) int {
    return a + b
}
```

### 📄 `math_test.go`

```go
package math

import "testing"

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    if result != 5 {
        t.Errorf("Expected 5, got %d", result)
    }
}
```

Laufen lassen:

```bash
go test
```

✅ ✅

---

## 🧪 Beispiel 2: Mit `testify` Assertions

Installieren:

```bash
go get github.com/stretchr/testify
```

Dann:

```go
import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
    assert.Equal(t, 5, Add(2, 3))
}
```

✔️ Kürzer, lesbarer

---

## 🧪 Beispiel 3: Tests mit `table-driven` style

```go
func TestAddTableDriven(t *testing.T) {
    tests := []struct {
        name     string
        a, b     int
        expected int
    }{
        {"positive numbers", 1, 2, 3},
        {"zero", 0, 0, 0},
        {"negative numbers", -1, -2, -3},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            result := Add(tt.a, tt.b)
            if result != tt.expected {
                t.Errorf("Expected %d, got %d", tt.expected, result)
            }
        })
    }
}
```

📋 Sauber, erweiterbar, ideal für viele Eingaben.

---

## 🧪 Beispiel 4: Funktion mit Abhängigkeit (Mock)

Angenommen du hast:

```go
func GetUserName(api APIClient, id int) string {
    return api.FetchName(id)
}
```

Dann kannst du im Test einen **MockClient** bauen:

```go
type mockAPI struct{}

func (m *mockAPI) FetchName(id int) string {
    return "TestUser"
}
```

Und im Test:

```go
func TestGetUserName(t *testing.T) {
    client := &mockAPI{}
    name := GetUserName(client, 42)
    assert.Equal(t, "TestUser", name)
}
```

---

## 🔥 Bonus: Coverage prüfen

```bash
go test -cover
```

Oder HTML-Report:

```bash
go test -coverprofile=coverage.out
go tool cover -html=coverage.out
```

Siehst genau, was getestet ist und was nicht!

---

## 🧠 Best Practices

✅ Testdateien enden auf `_test.go`  
✅ Benutze `t.Run` für Subtests  
✅ Schreibe kleine, fokussierte Tests  
✅ Mocks statt echte Datenbank/HTTP  
✅ Nutze `testify` für saubere Assertions

---

## 🚀 Willst du...

- ✅ Datenbank- oder REST-Handler testen?
- ✅ Tests mit Setup/Teardown?
- ✅ gRPC-Tests mit Mocks?

Ich helf dir dabei! Sag einfach, was du brauchst 😎