# 🔗 URL Shortener

A production-ready, containerized URL shortening service built with **Go**, powered by **PostgreSQL**, **Redis**, and **Kafka**. Designed for performance, scalability, and observability.

---

## 🛠️ Tech Stack

- **Go** — Backend logic
- **PostgreSQL** — Persistent URL store
- **Redis** — Caching layer
- **Kafka** — Visit logging and streaming
- **Docker Compose** — Service orchestration

---

## 🚀 Features

- ✅ Shorten URLs with optional custom aliases
- ✅ Redirect from short URL to original
- ✅ Redis caching for fast lookups
- ✅ Kafka stream for visit logs (IP + timestamp)
- ✅ PostgreSQL auto-initialized via SQL script

---

## 📦 Getting Started

### 1. Clone the repository and enter the directory

```bash
git clone https://github.com/your-username/url_shortener.git
cd url_shortener
```

### 2. Start the services

```bash
docker compose up --build
```

This starts:
- app (Go backend) at http://localhost:8080
- postgres on port 5433
- redis on port 6379
- kafka (with zookeeper) for logging events

> 📄 PostgreSQL is initialized using init.sql.

---

## 📡 API Endpoints

### 🔹 `POST /shorten`

Create a shortened URL.

**Request JSON:**

```json
{
  "url": "https://example.com",
  "customAlias": "ganesh"
}
```

**Curl Example:**
```bash
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com", "customAlias": "ganesh"}'
```

**Response:**
```json
{
  "shortUrl": "http://localhost:8080/ganesh"
}
```

### 🔹 `GET /{shortCode}`
Redirect to the original URL.

**Example:**
```bash
curl -v http://localhost:8080/ganesh
```

**Response:**
```
HTTP/1.1 302 Found
Location: https://example.com
```
---

## 🔁 Kafka Integration

Kafka is used to asynchronously log every URL visit.

### 🔸 Topic: `url_visits`

Each time a short URL is accessed (`GET /{shortCode}`), a message is published to the Kafka topic `url_visits` with the following format:

```
<timestamp> - visited from <ip>
```

**To consume messages from Kafka:**
```bash
docker exec -it <kafka-container-name> kafka-console-consumer \
  --bootstrap-server kafka:9092 \
  --topic url_visits --from-beginning
```

Replace <kafka-container-name> with the name or ID of your running Kafka container (e.g., url_shortener-kafka-1).

---

## 📂 Project Structure
```tree
.
├── db
│   ├── postgres.go
│   ├── redis.go
│   └── url.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── handlers
│   ├── redirect.go
│   └── shorten.go
├── init.sql
├── kafka
│   └── kafka.go
├── main.go
├── README.md
├── url_shortener
└── utils
    └── utils.go

```

---

## Cleanup
```bash
docker compose down -v
```

---