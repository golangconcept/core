Klar! 😎 Ich zeig dir Schritt für Schritt, wie du in **Go** Daten aus **MongoDB** abrufen kannst – mit einem praktischen Beispiel.

---

## 🧰 Voraussetzungen

1. Du hast MongoDB laufen (z. B. lokal oder Atlas).
2. Du hast Go installiert.
3. Du brauchst das MongoDB Go-Driver-Paket:

```bash
go get go.mongodb.org/mongo-driver/mongo
go get go.mongodb.org/mongo-driver/mongo/options
```

---

## 📦 Beispiel: Benutzer aus MongoDB abrufen

### 📂 Projektstruktur

```
meinprojekt/
├── main.go
└── go.mod
```

---

### 🛠 `go.mod`

```go
module meinprojekt

go 1.21
```

---

### 📄 `main.go`

```go
package main

import (
    "context"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

// Beispiel-Datenstruktur
type Benutzer struct {
    Name  string `bson:"name"`
    Email string `bson:"email"`
}

func main() {
    // Kontext mit Timeout (Verbindung abbrechen, wenn's zu lange dauert)
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    // MongoDB-Verbindung (Standard-Port)
    client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
    if err != nil {
        log.Fatal(err)
    }
    defer client.Disconnect(ctx)

    // Zugriff auf DB und Collection
    collection := client.Database("meinshop").Collection("benutzer")

    // 🔍 Alle Benutzer abrufen
    cursor, err := collection.Find(ctx, bson.M{})
    if err != nil {
        log.Fatal(err)
    }
    defer cursor.Close(ctx)

    fmt.Println("Benutzer aus MongoDB:")
    for cursor.Next(ctx) {
        var user Benutzer
        if err := cursor.Decode(&user); err != nil {
            log.Println("Fehler beim Dekodieren:", err)
            continue
        }
        fmt.Printf("- %s (%s)\n", user.Name, user.Email)
    }

    if err := cursor.Err(); err != nil {
        log.Fatal(err)
    }
}
```

---

## 📋 Beispiel-Daten in MongoDB (zum Testen)

In der MongoDB Shell oder MongoDB Compass:

```json
{
  "name": "Alice",
  "email": "alice@example.com"
}
```

---

## ▶️ Ausführen

```bash
go run main.go
```

### ✅ Beispielausgabe:

```
Benutzer aus MongoDB:
- Alice (alice@example.com)
```

---

## 🔁 Du willst mehr?

- ✅ Nach bestimmten Nutzern filtern (`FindOne`)
- ✅ Neue Benutzer einfügen (`InsertOne`)
- ✅ REST API drumrum bauen
- ✅ Verbindung zu MongoDB Atlas

Sag einfach, was du brauchst – ich helf dir beim nächsten Schritt.