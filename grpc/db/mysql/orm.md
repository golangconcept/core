Perfekt! Wenn du **MySQL mit ORM in Go** nutzen willst, ist der beliebteste und stabilste Weg:

> 👉 **[GORM](https://gorm.io/)** – eine mächtige ORM-Bibliothek für Go.

---

## ✅ Was du mit GORM bekommst

- Einfache Models (Structs = Tabellen)
- Automatische Migrationen
- Einfache Queries (lesen, schreiben, filtern)
- Unterstützung für MySQL, Postgres, SQLite, etc.

---

## 🚀 Schritt-für-Schritt: GORM + MySQL in Go

### 🔧 1. 📦 GORM & MySQL-Driver installieren

```bash
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

---

### 📄 2. Model erstellen

```go
// models/product.go
package models

type Product struct {
    ID    uint   `gorm:"primaryKey"`
    Name  string
    Price float64
}
```

---

### 🧠 3. Setup-Funktion für GORM mit MySQL

```go
// db/mysql.go
package db

import (
    "fmt"
    "log"
    "os"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "myapp/models"
)

func SetupMySQL() *gorm.DB {
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

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        user, pass, host, port, dbname)

    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("❌ Verbindung zu MySQL fehlgeschlagen: %v", err)
    }

    // Automatische Migration
    db.AutoMigrate(&models.Product{})

    fmt.Println("✅ MySQL mit GORM verbunden!")
    return db
}
```

---

### 🚀 4. Verwenden in `main.go`

```go
package main

import (
    "fmt"
    "myapp/db"
    "myapp/models"
)

func main() {
    database := db.SetupMySQL()

    // ✅ Neues Produkt erstellen
    newProduct := models.Product{Name: "Tisch", Price: 99.99}
    database.Create(&newProduct)

    // ✅ Alle Produkte abrufen
    var products []models.Product
    database.Find(&products)

    for _, p := range products {
        fmt.Printf("📦 Produkt: %s (%.2f€)\n", p.Name, p.Price)
    }
}
```

---

## 🌱 `.env` Datei (optional)

```env
MYSQL_USER=root
MYSQL_PASS=123456
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_DB=mydb
```

> Lade sie mit `github.com/joho/godotenv`, falls gewünscht.

---

## 🛠 GORM Quick Cheatsheet

| Aktion         | Code                                  |
|----------------|----------------------------------------|
| Einfügen       | `db.Create(&obj)`                      |
| Alle holen     | `db.Find(&slice)`                      |
| Filtern        | `db.Where("name = ?", "Tisch").Find()` |
| Aktualisieren  | `db.Model(&obj).Update("price", 10)`   |
| Löschen        | `db.Delete(&obj)`                      |

---

## ✅ TL;DR

- GORM ist top für saubere, strukturierte MySQL-Nutzung in Go
- Automatische Migrationsfunktion inklusive
- Nutzt `structs` als Models (ORM-Style)
- Super für CRUD-Backends & REST APIs

---

Willst du ein REST API mit GORM + Gin bauen? Oder Auth, Pagination, oder Testdaten einfügen? Sag einfach, was dein Ziel ist!