Klar! Lass uns deine `NewMySQL()`-Funktion **sauberer, sicherer und konfigurierbarer** machen – ganz im Stil von **Clean Code** und **Best Practices** ✅

---

## ✅ Verbesserte `NewMySQL()` mit:

- Umgebungsvariablen (statt Hardcoding)
- Timeout + Connection Pooling
- Fehlerprüfung durch `Ping()`
- Konfigurierbare DSN-Logik
- Logging

---

### 📁 Datei: `pkg/db/mysql.go` (standardisiertes Setup)

```go
package db

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    "time"

    _ "github.com/go-sql-driver/mysql"
)

type Config struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    Timeout  time.Duration
}

func NewMySQL(cfg Config) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&timeout=%s",
        cfg.User,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.DBName,
        cfg.Timeout.String(),
    )

    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("could not open MySQL: %w", err)
    }

    // Pool settings (optional tuning)
    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(30 * time.Minute)

    // Check connection
    if err := db.Ping(); err != nil {
        return nil, fmt.Errorf("could not connect to MySQL: %w", err)
    }

    log.Println("✅ MySQL connected successfully")
    return db, nil
}
```

---

### 📁 Datei: `cmd/server/main.go` (Verwendung mit Env)

```go
mysqlCfg := db.Config{
    Host:     os.Getenv("MYSQL_HOST"),
    Port:     os.Getenv("MYSQL_PORT"),
    User:     os.Getenv("MYSQL_USER"),
    Password: os.Getenv("MYSQL_PASSWORD"),
    DBName:   os.Getenv("MYSQL_DB"),
    Timeout:  5 * time.Second,
}

dbConn, err := db.NewMySQL(mysqlCfg)
if err != nil {
    log.Fatalf("❌ Failed to connect MySQL: %v", err)
}
```

---

## 📦 `.env` (optional oder via `os.Getenv` direkt im Compose)

```
MYSQL_HOST=mysql
MYSQL_PORT=3306
MYSQL_USER=user
MYSQL_PASSWORD=password
MYSQL_DB=yourdb
```

> Kannst du mit [`github.com/joho/godotenv`](https://github.com/joho/godotenv) automatisch laden, wenn du lokal entwickelst.

```go
_ = godotenv.Load(".env")
```

---

## 🧪 Vorteile dieser Version:

✅ Klar getrennte **Konfiguration**  
✅ Gute **Fehlerbehandlung & Logging**  
✅ Produktionstauglich mit **Timeouts & Pooling**  
✅ Einfach testbar (via Dependency Injection)  

---

Wenn du willst, kann ich dir:

- Eine komplette `pkg/config`-Struktur generieren  
- Optional ENV-Validator  
- YAML/JSON-Ladefunktion für Settings

Sag einfach: **"Mach mir die config-Komponente dazu"** oder **"Zeig mit ENV-Loader + Validator"**, und ich baue das gleich mit rein 💡