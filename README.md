# test-go

A Go backend authentication service with Register, Login, and JWT-protected APIs.

## Tech Stack

| Layer | Package |
|---|---|
| HTTP Framework | [Gin](https://github.com/gin-gonic/gin) |
| ORM | [GORM](https://gorm.io) |
| Database | PostgreSQL |
| Auth | [golang-jwt/jwt v5](https://github.com/golang-jwt/jwt) |
| Password | bcrypt (`golang.org/x/crypto`) |
| Config | [Viper](https://github.com/spf13/viper) + godotenv |
| Logging | [Uber Zap](https://github.com/uber-go/zap) |

## Project Structure

```
test-go/
├── cmd/server/main.go               # Entry point
├── internal/
│   ├── config/config.go             # Config loader
│   ├── database/database.go         # GORM connection
│   ├── middleware/auth_middleware.go # JWT guard
│   ├── model/user.go                # User model
│   ├── handler/auth_handler.go      # HTTP handlers
│   ├── service/auth_service.go      # Business logic
│   ├── repository/user_repository.go# DB queries
│   └── utils/
│       ├── jwt.go                   # JWT generate/validate
│       └── password.go              # bcrypt helpers
├── docker-compose.yml               # PostgreSQL
├── .env.example
└── go.mod
```

## Quick Start

### 1. Start PostgreSQL

```bash
docker compose up -d
```

### 2. Configure environment

```bash
cp .env.example .env
# Edit .env if needed
```

### 3. Install dependencies & run

```bash
go mod tidy
go run ./cmd/server
```

Server starts on `http://localhost:8080`

## API Endpoints

| Method | Path | Auth | Description |
|---|---|---|---|
| `GET` | `/health` | ❌ | Health check |
| `POST` | `/api/v1/auth/register` | ❌ | Register a new user |
| `POST` | `/api/v1/auth/login` | ❌ | Login, get JWT token |
| `GET` | `/api/v1/auth/me` | ✅ Bearer | Get current user info |

## Example Requests

```bash
# Register
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"name":"Doni","email":"doni@test.com","password":"pass123"}'

# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"doni@test.com","password":"pass123"}'

# Get current user (replace TOKEN)
curl http://localhost:8080/api/v1/auth/me \
  -H "Authorization: Bearer TOKEN"
```
