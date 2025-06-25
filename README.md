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

## Test Results for 10K Virtual Users
```


         /\      Grafana   /‾‾/  
    /\  /  \     |\  __   /  /   
   /  \/    \    | |/ /  /   ‾‾\ 
  /          \   |   (  |  (‾)  |
 / __________ \  |_|\_\  \_____/ 

     execution: local
        script: load-test.js
        output: -

     scenarios: (100.00%) 1 scenario, 1000 max VUs, 1m0s max duration (incl. graceful stop):
              * default: 1000 looping VUs for 30s (gracefulStop: 30s)

WARN[0011] The test has generated metrics with 100080 unique time series, which is higher than the suggested limit of 100000 and could cause high memory usage. Consider not using high-cardinality values like unique IDs as metric tags or, if you need them in the URL, use the name metric tag or URL grouping. See https://grafana.com/docs/k6/latest/using-k6/tags-and-groups/ for details.  component=metrics-engine-ingester
WARN[0023] The test has generated metrics with 200637 unique time series, which is higher than the suggested limit of 100000 and could cause high memory usage. Consider not using high-cardinality values like unique IDs as metric tags or, if you need them in the URL, use the name metric tag or URL grouping. See https://grafana.com/docs/k6/latest/using-k6/tags-and-groups/ for details.  component=metrics-engine-ingester


  █ TOTAL RESULTS 

    checks_total.......................: 87354   2828.610657/s
    checks_succeeded...................: 100.00% 87354 out of 87354
    checks_failed......................: 0.00%   0 out of 87354

    ✓ Shorten status was 200
    ✓ Redirect status was 302
    ✓ Redirect location correct

    HTTP
    http_req_duration.......................................................: avg=24.23ms min=139.67µs med=1.52ms max=2.52s p(90)=14.51ms p(95)=26.71ms
      { expected_response:true }............................................: avg=24.23ms min=139.67µs med=1.52ms max=2.52s p(90)=14.51ms p(95)=26.71ms
    http_req_failed.........................................................: 0.00%  0 out of 58236
    http_reqs...............................................................: 58236  1885.740438/s

    EXECUTION
    iteration_duration......................................................: avg=1.04s   min=1s       med=1s     max=3.56s p(90)=1.02s   p(95)=1.06s  
    iterations..............................................................: 29118  942.870219/s
    vus.....................................................................: 1000   min=1000       max=1000
    vus_max.................................................................: 1000   min=1000       max=1000

    NETWORK
    data_received...........................................................: 11 MB  344 kB/s
    data_sent...............................................................: 8.7 MB 282 kB/s




running (0m30.9s), 0000/1000 VUs, 29118 complete and 0 interrupted iterations
default ✓ [======================================] 1000 VUs  30s

```