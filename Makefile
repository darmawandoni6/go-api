# ============================================================================
#  test-go — Makefile
# ============================================================================

APP_NAME   := test-go
CMD_PATH   := ./cmd/server
BIN_DIR    := ./bin
BINARY     := $(BIN_DIR)/$(APP_NAME)
GO         := go
GOFLAGS    := -ldflags="-s -w"   # strip debug info for smaller binary

.PHONY: help run build clean tidy lint test \
        db-up docker-up docker-down docker-build docker-logs docker-restart

# Default target — show help
.DEFAULT_GOAL := help

# ── Help ─────────────────────────────────────────────────────────────────────
help:
	@echo ""
	@echo "  $(APP_NAME) — available commands"
	@echo ""
	@echo "  Dev"
	@echo "    make run          Run the server locally (hot .env reload)"
	@echo "    make build        Compile binary to $(BINARY)"
	@echo "    make clean        Remove compiled binaries"
	@echo "    make tidy         Tidy and verify go.mod / go.sum"
	@echo "    make lint         Run go vet static analysis"
	@echo "    make test         Run all unit tests"
	@echo ""
	@echo "  Docker / Podman"
	@echo "    make db-up        Start only PostgreSQL (for local 'make run')"
	@echo "    make docker-up    Start ALL containers (postgres + app)"
	@echo "    make docker-down  Stop and remove containers"
	@echo "    make docker-build Rebuild app image"
	@echo "    make docker-logs  Tail logs from all containers"
	@echo "    make docker-restart  Restart the app container only"
	@echo ""

# ── Development ───────────────────────────────────────────────────────────────
run:
	$(GO) run $(CMD_PATH)

build:
	@mkdir -p $(BIN_DIR)
	$(GO) build $(GOFLAGS) -o $(BINARY) $(CMD_PATH)
	@echo "✅  Binary built: $(BINARY)"

clean:
	@rm -rf $(BIN_DIR)
	@echo "🗑️   Cleaned $(BIN_DIR)"

tidy:
	$(GO) mod tidy
	$(GO) mod verify
	@echo "✅  go.mod is tidy"

lint:
	$(GO) vet ./...
	@echo "✅  go vet passed"

test:
	$(GO) test ./... -v -count=1

# ── Docker / Podman ───────────────────────────────────────────────────────────
db-up:
	podman compose up -d postgres

docker-up:
	podman compose up -d

docker-down:
	podman compose down

docker-build:
	podman compose build app

docker-logs:
	podman compose logs -f

docker-restart:
	podman compose restart app
