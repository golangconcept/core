Yes! Let's **Dockerize** the entire Go + gRPC + REST API setup so it runs cleanly inside containers — and optionally connect it to your Nuxt frontend.

---

## 📦 Dockerized Setup Overview

We'll containerize:

1. 🐹 Go server (gRPC + REST API)
2. 🌐 Nuxt frontend (optional)
3. 🧾 Use Docker Compose to run everything

---

## 🗂️ Project Structure (Final Form)

```
go-grpc-example/
├── proto/
│   └── product.proto
├── pb/
│   └── (generated files)
├── server/
│   ├── main.go
│   └── product_server.go
├── api/
│   └── main.go         # REST API server
├── Dockerfile          # for Go backend
├── docker-compose.yml
└── go.mod / go.sum
```

---

## 🐳 Step 1: Dockerfile (Go Backend)

Create a `Dockerfile` in the project root:

```Dockerfile
# syntax=docker/dockerfile:1

# 1. Build stage
FROM golang:1.20 AS builder

WORKDIR /app

# Copy go.mod files and download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy all source
COPY . .

# Build REST + gRPC server
RUN go build -o /grpc-server ./api/main.go

# 2. Run stage
FROM debian:bullseye-slim

COPY --from=builder /grpc-server /grpc-server

EXPOSE 8080 50051

ENTRYPOINT ["/grpc-server"]
```

This builds and runs your Go app as a single binary container.

---

## 🛠️ Step 2: docker-compose.yml

```yaml
version: "3.9"

services:
  backend:
    build: .
    ports:
      - "8080:8080"   # REST API
      - "50051:50051" # gRPC
    container_name: go-grpc-api
```

> You can later add Nuxt in the same file (I’ll show that too).

---

## 🐾 Step 3: Build & Run

In the project root:

```bash
docker-compose build
docker-compose up
```

---

## ✅ Test It

Once up and running:

- gRPC server is on `localhost:50051`
- REST API is on `localhost:8080`

Test in browser:

```
http://localhost:8080/api/products
```

---

## 🔌 Optional: Add Nuxt App to Compose

If your Nuxt frontend lives in `nuxt/`, update `docker-compose.yml`:

```yaml
  frontend:
    build:
      context: ./nuxt
    ports:
      - "3000:3000"
    container_name: nuxt-frontend
```

And create a `Dockerfile` inside `nuxt/`:

```Dockerfile
FROM node:18

WORKDIR /app

COPY . .

RUN npm install
RUN npm run build

EXPOSE 3000

CMD ["npm", "run", "start"]
```

---

## 💡 Dev Tips

### Rebuild after `.proto` change?

Regenerate code locally, or add it into your Dockerfile with `protoc`.

Or, use [buf.build](https://buf.build) for managed protobuf pipelines.

---

## ✅ Recap: You Now Have

- gRPC + REST Go API running in Docker
- REST endpoint exposed on port 8080
- Easy to integrate Nuxt or any frontend

---

Want me to:
- zip this whole structure up for you?
- generate a Nuxt sample frontend for testing this?

Let me know 💪