# 🌸 BloomService

**BloomService** is a high-performance, RESTful microservice written in Go that provides fast, memory-efficient **Bloom Filter** operations for key existence checks. Ideal for deduplication, caching, and spam filtering use cases in distributed systems.

## ✨ Features

- 🧠 In-memory Bloom Filter with configurable false-positive probability
- ⚡ Fast key insert and existence check
- 📊 Stats endpoint for runtime filter metrics
- 🛡️ RESTful API with input validation and clear status codes
- 📈 Prometheus metrics for observability
- 🐳 Dockerized with CI/CD support

## 📦 Use Cases

- Prevent duplicate processing in event streams
- API cache validation
- Spam and abuse detection
- Lightweight deduplication layer (e.g., email signups, IP checks)

## 🚀 API Endpoints

### ➕ Add a key
```http
POST /bloom/add
Content-Type: application/json

{
  "key": "vinit@example.com"
}
````

### ❓ Check a key

```http
GET /bloom/check?key=vinit@example.com
```

Returns:

```json
{
  "mightExist": true
}
```

### 📊 Stats

```http
GET /bloom/stats
```

Returns:

```json
{
  "insertions": 1000,
  "falsePositiveRate": 0.01,
  "capacity": 100000
}
```

## ⚙️ Configuration

| ENV Variable | Description                            | Default  |
| ------------ | -------------------------------------- | -------- |
| `PORT`       | Port to serve HTTP API                 | `8080`   |
| `CAPACITY`   | Expected number of elements            | `100000` |
| `FPP`        | False positive probability (0.01 = 1%) | `0.01`   |

## 🛠️ Development

```bash
# Clone and run
git clone https://github.com/vinit-chauhan/bloomservice.git
cd bloomservice
go run main.go
```

```bash
# Run tests
go test ./...
```

## 🐳 Docker

```bash
docker build -t bloomservice .
docker run -p 8080:8080 bloomservice
```

## 📈 Observability

BloomService exposes Prometheus metrics at `/metrics`:

* `bloomservice_requests_total`
* `bloomservice_errors_total`
* `bloomservice_false_positive_rate`

## 🧪 Roadmap

* [ ] Custom disk-based persistence
* [ ] gRPC interface alongside REST
* [ ] API authentication support
* [ ] Bloom filter expiration / TTL support
* [ ] Horizontal sharding for distributed usage

## 📄 License
MIT License. Contributions welcome!

## 👨‍💻 Author
Built with ❤️ by [Vinit Chauhan](https://github.com/vinit-chauhan) – Golang | Elasticsearch | Distributed Systems.
