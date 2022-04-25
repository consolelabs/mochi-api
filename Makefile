APP_NAME=mochi-api
DEFAULT_PORT=8200
.PHONY: setup init build dev test migrate-up migrate-down

setup:
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/vektra/mockery/v2@latest
	cp .env.sample .env
	make init

init: 
	make remove-infras
	docker-compose up -d
	@echo "Waiting for database connection..."
	@while ! docker exec mochi-postgres pg_isready > /dev/null; do \
		sleep 1; \
	done

remove-infras:
	docker-compose down --remove-orphans

build:
	env GOOS=darwin GOARCH=amd64 go build -o bin ./...

dev:
	go run ./cmd/server/main.go

test:
	@PROJECT_PATH=$(shell pwd) go test -cover ./...

migrate-new:
	sql-migrate new -env=local ${name}

migrate-up:
	sql-migrate up -env=local

migrate-down:
	sql-migrate down -env=local

docker-build:
	docker build \
	--build-arg DEFAULT_PORT="${DEFAULT_PORT}" \
	-t ${APP_NAME}:latest .

seed-db:
	@docker exec -t mochi-postgres sh -c "mkdir -p /seed"
	@docker exec -t mochi-postgres sh -c "rm -rf /seed/*"
	@docker cp migrations/seed mochi-postgres:/
	@docker exec -t mochi-postgres sh -c "PGPASSWORD=postgres psql -U postgres -d mochi_local -f /seed/seed.sql"

reset-db:
	make init
	make migrate-up
	make seed-db

init-test-db:
	go install github.com/rubenv/sql-migrate/...@latest
	make init
	make migrate-up
	make seed-db

