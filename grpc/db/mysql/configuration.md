Nice! Du willst stattdessen **MySQL** in deinem Go-Projekt verwenden? Kein Problem – ich zeig dir, wie du eine **robuste, standardisierte `setupMySQL`-Funktion** schreibst, ganz ähnlich wie bei MongoDB oder Redis.

---

## ✅ Ziel: MySQL-Verbindung einrichten (robust & flexibel)

Wir verwenden:

- `go-sql-driver/mysql` (Standard-Driver)
- `.env` / `os.Getenv()` für Konfiguration
- Ping zur Verbindungsprüfung
- `sql.DB` mit Verbindungs-Pooling

---

## 📦 Installation

```bash
go get github.com/go-sql-driver/mysql
```

---

## 🛠 Beispiel: `setupMySQL` Funktion

```go
package main

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

func setupMySQL() *sql.DB {
    // Umgebungsvariablen verwenden oder Fallback
    user := os.Getenv("MYSQL_USER")
    pass := os.Getenv("MYSQL_PASS")
    host := os.Getenv("MYSQL_HOST")
    port := os.Getenv("MYSQL_PORT")
    dbname := os.Getenv("MYSQL_DB")

    if user == "" {
        user = "root"
    }
    if pass == "" {
        pass = ""
    }
    if host == "" {
        host = "localhost"
    }
    if port == "" {
        port = "3306"
    }
    if dbname == "" {
        dbname = "mydb"
    }

    // Verbindungs-String
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, pass, host, port, dbname)

    // DB öffnen
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("❌ Fehler beim Öffnen der MySQL-Verbindung: %v", err)
    }

    // Optional: Verbindungspool konfigurieren
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(30 * time.Minute)

    // Verbindung testen
    if err := db.Ping(); err != nil {
        log.Fatalf("❌ Verbindung zu MySQL fehlgeschlagen: %v", err)
    }

    fmt.Println("✅ Verbunden mit MySQL!")
    return db
}
```

---

## 🧪 Verwendung in `main.go`

```go
func main() {
    db := setupMySQL()

    // Beispiel: Query
    rows, err := db.Query("SELECT id, name FROM products")
    if err != nil {
        log.Fatalf("Fehler bei SELECT: %v", err)
    }
    defer rows.Close()

    for rows.Next() {
        var id int
        var name string
        rows.Scan(&id, &name)
        fmt.Printf("- [%d] %s\n", id, name)
    }
}
```

---

## 🌱 `.env` Beispiel (optional mit `godotenv`)

```env
MYSQL_USER=root
MYSQL_PASS=passwort123
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_DB=mydb
```

Lade sie dann mit:

```go
import "github.com/joho/godotenv"

func init() {
    _ = godotenv.Load()
}
```

---

## ✅ TL;DR

| Was | Warum |
|-----|-------|
| `sql.Open(...)` | Öffnet und verwaltet die Verbindung |
| `Ping()` | Verbindung testen |
| `parseTime=true` | Damit `DATETIME` als `time.Time` behandelt wird |
| `SetMaxOpenConns()` etc. | Steuert Performance & Pooling |

---

Wenn du willst, zeig ich dir auch:

- Wie man **INSERT/UPDATE/DELETE** macht,
- wie man **SQL-Injection sicher verhindert** (mit Prepared Statements),
- oder ein **Repository Pattern** für größere Projekte baut.

Willst du mit `sqlc`, `gorm` oder `ent` ORM arbeiten – oder lieber "raw SQL"?