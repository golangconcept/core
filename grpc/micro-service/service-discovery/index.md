Nice! Let’s take your **Go microservice architecture** to the next level with:

### ✅ **Monitoring** (Prometheus + Grafana)  
### ✅ **Service Discovery** (Consul)

This setup is **real production-grade** — and yes, it’s fully Dockerized 💥

---

## 🧱 Final Architecture Overview

```
go-microservices/
├── category-service/           # gRPC + Mongo
├── product-service/            # gRPC + Mongo + Redis
├── proto/                      # Shared .proto files
├── api-gateway/ (optional)     # REST/Nuxt interface
├── consul/                     # Service discovery agent config
├── prometheus/                 # Prometheus config
├── grafana/                    # Dashboards
├── docker-compose.yml
```

---

## 🔧 SERVICE DISCOVERY — using Consul

### 🗂️ In each service (e.g. `product-service`):

Register service with Consul on startup:

```go
func registerWithConsul() {
    config := api.DefaultConfig()
    config.Address = "http://consul:8500"

    client, _ := api.NewClient(config)
    registration := &api.AgentServiceRegistration{
        Name:    "product-service",
        Address: "product-service",
        Port:    50052,
        Check: &api.AgentServiceCheck{
            GRPC:     "product-service:50052",
            Interval: "10s",
            Timeout:  "1s",
        },
    }
    _ = client.Agent().ServiceRegister(registration)
}
```

Call this during `main()` startup.

Now any service (e.g. `product-service`) can **discover `category-service`** from Consul, not hardcoded IPs.

---

## 📈 MONITORING — Prometheus + Grafana

### Step 1: Add Prometheus metrics to your Go services

Use [`prometheus/client_golang`](https://github.com/prometheus/client_golang):

```go
import "github.com/prometheus/client_golang/prometheus/promhttp"

go func() {
    http.Handle("/metrics", promhttp.Handler())
    log.Fatal(http.ListenAndServe(":2112", nil))
}()
```

> Run this alongside gRPC to expose `/metrics` endpoint.

---

### Step 2: `prometheus/prometheus.yml`

```yaml
global:
  scrape_interval: 5s

scrape_configs:
  - job_name: 'category-service'
    static_configs:
      - targets: ['category-service:2112']
  - job_name: 'product-service'
    static_configs:
      - targets: ['product-service:2112']
```

---

### Step 3: Add Grafana

#### `docker-compose.yml` additions:

```yaml
services:
  prometheus:
    image: prom/prometheus
    volumes:
      - ./prometheus/prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"

  grafana:
    image: grafana/grafana
    ports:
      - "3000:3000"
    depends_on:
      - prometheus
```

---

### 🎨 Grafana Dashboards

- Open Grafana `localhost:3000` (admin / admin)
- Add Prometheus as data source (`http://prometheus:9090`)
- Import dashboard: "Go gRPC server metrics"

---

## 🧠 Optional: Auto Service Discovery via gRPC Resolver

Use a **Consul gRPC resolver** package (like [`grpc-consul-resolver`](https://github.com/mbobakov/grpc-consul-resolver)):

```go
import _ "github.com/mbobakov/grpc-consul-resolver"

conn, err := grpc.Dial(
    "consul://category-service:8500/category-service?wait=14s",
    grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
    grpc.WithInsecure(),
)
```

> No need to hardcode IPs anymore. Auto-balancing included.

---

## ✅ Final Checklist

| Component         | Tech Used                  |
|------------------|----------------------------|
| gRPC transport    | `google.golang.org/grpc`   |
| Service discovery | `Consul` + gRPC resolver   |
| Monitoring        | `Prometheus`, `Grafana`    |
| Database          | MongoDB (per service)      |
| Caching           | Redis (product service)    |
| Metrics           | `/metrics` on `:2112`      |

---

## ✅ Want this as a template?

I can generate:
- ✅ Category + Product service
- ✅ Shared Protobuf + Registry
- ✅ Fully Dockerized with Prometheus, Grafana, Consul
- ✅ Makefile / bash scripts to run everything

Say: **“Create full microservice template”**, and I’ll prepare it for you 🔥

Or, if you want to start with just **category + product with monitoring**, I can do that too.

Let me know how you'd like to proceed 💬