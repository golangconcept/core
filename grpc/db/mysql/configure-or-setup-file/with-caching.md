Sehr cool – wir bauen jetzt eine **standardisierte MySQL + Redis-Anbindung** in Go im Clean-Code-Stil.  
💡 Perfekt für Microservices, REST oder gRPC.

---

## ✅ Ziel:  
Eine strukturierte, produktionsreife Initialisierung für:

- `MySQL` (mit Connection-Pooling, ENV-Config, Ping)
- `Redis` (mit Timeout, ENV-Config, Ping)

---

## 📁 Strukturvorschlag

```
server/
├── cmd/
│   └── main.go
├── pkg/
│   ├── config/           # ENV-Loader
│   ├── db/               # MySQL Initializer
│   └── cache/            # Redis Initializer
```

---

## 📦 `pkg/config/config.go` – Umgebungsvariablen laden

```go
package config

import (
    "os"
    "time"
)

type MySQLConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    Timeout  time.Duration
}

type RedisConfig struct {
    Addr     string
    Password string
    DB       int
}

type AppConfig struct {
    MySQL MySQLConfig
    Redis RedisConfig
}

func Load() AppConfig {
    return AppConfig{
        MySQL: MySQLConfig{
            Host:     getEnv("MYSQL_HOST", "localhost"),
            Port:     getEnv("MYSQL_PORT", "3306"),
            User:     getEnv("MYSQL_USER", "user"),
            Password: getEnv("MYSQL_PASSWORD", "password"),
            DBName:   getEnv("MYSQL_DB", "mydb"),
            Timeout:  5 * time.Second,
        },
        Redis: RedisConfig{
            Addr:     getEnv("REDIS_ADDR", "localhost:6379"),
            Password: getEnv("REDIS_PASSWORD", ""),
            DB:       0,
        },
    }
}

func getEnv(key, fallback string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return fallback
}
```

---

## 🐬 `pkg/db/mysql.go` – MySQL-Setup

```go
package db

import (
    "database/sql"
    "fmt"
    "time"

    _ "github.com/go-sql-driver/mysql"
    "yourproject/pkg/config"
)

func NewMySQL(cfg config.MySQLConfig) (*sql.DB, error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&timeout=%s",
        cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName, cfg.Timeout.String(),
    )
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        return nil, err
    }

    db.SetMaxOpenConns(10)
    db.SetMaxIdleConns(5)
    db.SetConnMaxLifetime(30 * time.Minute)

    if err := db.Ping(); err != nil {
        return nil, err
    }

    return db, nil
}
```

---

## 🧠 `pkg/cache/redis.go` – Redis-Setup

```go
package cache

import (
    "context"
    "time"

    "github.com/redis/go-redis/v9"
    "yourproject/pkg/config"
)

func NewRedis(cfg config.RedisConfig) (*redis.Client, error) {
    client := redis.NewClient(&redis.Options{
        Addr:     cfg.Addr,
        Password: cfg.Password,
        DB:       cfg.DB,
    })

    ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
    defer cancel()

    if err := client.Ping(ctx).Err(); err != nil {
        return nil, err
    }

    return client, nil
}
```

---

## 🚀 `cmd/main.go` – Zusammensetzen

```go
package main

import (
    "log"

    "yourproject/pkg/cache"
    "yourproject/pkg/config"
    "yourproject/pkg/db"
)

func main() {
    cfg := config.Load()

    mysqlDB, err := db.NewMySQL(cfg.MySQL)
    if err != nil {
        log.Fatalf("❌ MySQL error: %v", err)
    }
    defer mysqlDB.Close()

    redisClient, err := cache.NewRedis(cfg.Redis)
    if err != nil {
        log.Fatalf("❌ Redis error: %v", err)
    }
    defer redisClient.Close()

    log.Println("✅ DB + Redis init erfolgreich!")
    // Hier kannst du deine Services starten (gRPC, REST etc.)
}
```

---

## 🧪 `.env` Beispiel

```
MYSQL_HOST=localhost
MYSQL_PORT=3306
MYSQL_USER=root
MYSQL_PASSWORD=secret
MYSQL_DB=shopdb

REDIS_ADDR=localhost:6379
REDIS_PASSWORD=
```

---

## ✅ Vorteile dieses Setups

| Feature           | Beschreibung                        |
|------------------|-------------------------------------|
| Clean structure   | Trennung von Config, DB, Cache      |
| Modular           | Wiederverwendbar in allen Services  |
| Konfigurierbar    | Per ENV oder `.env`                 |
| Sicher            | Zeitüberschreitungen + Ping         |
| Production-ready  | Connection Pooling, Fallbacks       |

---

## ➕ Bonus: Caching nutzen

Wenn du willst, zeig ich dir z. B. wie du:

- ein `[]Product` Ergebnis in Redis cachen kannst  
- das Cache-Verhalten im gRPC Service implementierst  
- TTLs einstellst

Sag einfach:  
👉 **"Zeig mir Redis-Caching für Produktliste"**  
oder  
👉 **"Mach ein Beispiel mit Set/Get TTL"**

Ich helf dir direkt weiter 🚀