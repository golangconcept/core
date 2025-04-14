Super! Du willst jetzt:

✅ Eine **Daten-Fetching API in Go**  
✅ Und dazu eine **GraphQL-Schnittstelle**  

Ich zeig dir **beides in einem klaren Setup** – du bekommst:

- Eine einfache REST API mit `net/http` oder `Gin`  
- Eine GraphQL API mit `graphql-go` oder `gqlgen`  
- Optionale MongoDB-Datenquelle  
- Beispiel: Produkte & Kategorien

---

## 🚀 1. REST API (Beispiel mit `net/http`)

### ✨ Datei: `main.go`

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func getProducts(w http.ResponseWriter, r *http.Request) {
	products := []Product{
		{ID: 1, Name: "MacBook", Price: 2499},
		{ID: 2, Name: "iPhone", Price: 999},
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func main() {
	http.HandleFunc("/api/products", getProducts)
	log.Println("REST server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
```

Aufrufbar via: `http://localhost:8080/api/products`

---

## 🧠 2. GraphQL API (mit `gqlgen` empfohlen)

### 🔧 Setup gqlgen

1. Installiere gqlgen:
   ```bash
   go install github.com/99designs/gqlgen@latest
   ```

2. Initialisiere:
   ```bash
   go mod init graphql-api
   go run github.com/99designs/gqlgen init
   ```

3. Dann erstelle dein Schema:

### 📄 Datei: `graph/schema.graphqls`

```graphql
type Product {
  id: Int!
  name: String!
  price: Int!
}

type Query {
  products: [Product!]!
}
```

### ⚙️ Resolver: `graph/resolver.go`

```go
package graph

import "graphql-api/graph/model"

type Resolver struct{}

func (r *queryResolver) Products() ([]*model.Product, error) {
	return []*model.Product{
		{ID: 1, Name: "MacBook", Price: 2499},
		{ID: 2, Name: "iPhone", Price: 999},
	}, nil
}
```

> `model.Product` wird automatisch generiert

---

### 🚀 Starten

```bash
go run server.go
```

GraphQL Playground öffnet sich unter:  
👉 `http://localhost:8080`

Du kannst abfragen wie:

```graphql
query {
  products {
    id
    name
    price
  }
}
```

---

## 🔄 Bonus: MongoDB anbinden

Du kannst im Resolver `Products()` z.B. mit MongoDB verbinden:

```go
client, _ := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
collection := client.Database("shop").Collection("products")
cursor, _ := collection.Find(ctx, bson.M{})
```

---

## ✅ REST oder GraphQL?

| REST                  | GraphQL              |
|-----------------------|----------------------|
| Einfach zu implementieren | Flexibel für Clients     |
| Gute Tooling           | Typensystem & Playground |
| Besser für einfache APIs | Besser für komplexe Daten |

---

## 📦 Willst du ein komplettes Beispielprojekt?

Ich kann dir eine ZIP oder ein GitHub-Repo mit folgendem machen:

- REST + GraphQL parallel
- MongoDB-Anbindung
- Dockerfile + Compose
- Sample-Daten
- Nuxt-Frontend ready

👉 Sag einfach: **„Erstelle GraphQL + REST Projekt“**  
Oder: **„Mach’s Docker ready“**

Ich helf dir sofort weiter 😎