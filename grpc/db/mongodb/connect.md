Sehr gute Idee! 💪 Du willst deine `setupMongo`-Funktion robuster und standardisierter machen — also mit:

- **Konfigurierbarer URI (z. B. via Umgebungsvariable)**
- **Verbindungsprüfung (`Ping`)**
- **Timeout-Handling**
- **Logging**
- Fehler-Handling mit Graceful Exit (statt nur abrupt `log.Fatalf`)

---

## ✅ Verbessert & robuster: `setupMongo()`

```go
package main

import (
    "context"
    "fmt"
    "log"
    "os"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func setupMongo(ctx context.Context) *mongo.Collection {
    // URI über Umgebungsvariable laden oder Fallback verwenden
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        mongoURI = "mongodb://localhost:27017" // Standardwert
    }

    // Kontextspezifisches Timeout setzen
    ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
    defer cancel()

    clientOpts := options.Client().ApplyURI(mongoURI)
    client, err := mongo.Connect(ctx, clientOpts)
    if err != nil {
        log.Fatalf("❌ MongoDB-Verbindungsfehler: %v", err)
    }

    // Verbindung testen (Ping)
    if err := client.Ping(ctx, nil); err != nil {
        log.Fatalf("❌ MongoDB nicht erreichbar: %v", err)
    }

    fmt.Printf("✅ Verbunden mit MongoDB: %s\n", mongoURI)

    return client.Database("mydb").Collection("products")
}
```

---

## 🧪 Anwendung in `main.go`

```go
func main() {
    ctx := context.Background()
    collection := setupMongo(ctx)

    // Beispiel: Anzahl der Produkte zählen
    count, err := collection.CountDocuments(ctx, bson.M{})
    if err != nil {
        log.Fatal("Fehler beim Zählen der Dokumente:", err)
    }

    fmt.Println("📦 Anzahl Produkte:", count)
}
```

---

## 📦 Optional: `.env` verwenden

Falls du deine MongoDB-URI nicht hardcoden willst:

1. `.env` Datei anlegen:

```env
MONGO_URI=mongodb://mongo:27017
```

2. Mit `godotenv` laden:

```bash
go get github.com/joho/godotenv
```

3. In deinem `main.go`:

```go
import "github.com/joho/godotenv"

func init() {
    _ = godotenv.Load()
}
```

---

## ✅ TL;DR

| Feature                 | Vorteil                        |
|-------------------------|-------------------------------|
| `os.Getenv()`           | flexibel durch ENV-Variablen  |
| `Ping()`                | prüft, ob DB wirklich erreichbar |
| `context.WithTimeout()` | verhindert „hängende“ Verbindungen |
| `defer cancel()`        | saubere Ressourcenfreigabe     |

---

Wenn du magst, zeig ich dir gern noch, wie man:
- mehrere Collections managed,
- MongoDB-Daten cached,
- oder ein Repository Pattern draus macht.

Sag einfach, was dein Ziel ist! 🚀