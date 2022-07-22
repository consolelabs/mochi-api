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
	go install github.com/rubenv/sql-migrate/...@latest	
	make remove-infras
	docker-compose up -d
	@echo "Waiting for database connection..."
	@while ! docker exec mochi-postgres pg_isready > /dev/null; do \
		sleep 1; \
	done
	make migrate-up
	make seed-db

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

gen-mock:
	@mockgen -source=./pkg/repo/guild_user_xp/store.go -destination=./pkg/repo/guild_user_xp/mocks/store.go
	@mockgen -source=./pkg/repo/guild_user_activity_log/store.go -destination=./pkg/repo/guild_user_activity_log/mocks/store.go
	@mockgen -source=./pkg/repo/invite_histories/store.go -destination=./pkg/repo/invite_histories/mocks/store.go
	@mockgen -source=./pkg/repo/discord_guilds/store.go -destination=./pkg/repo/discord_guilds/mocks/store.go
	@mockgen -source=./pkg/repo/guild_custom_command/store.go -destination=./pkg/repo/guild_custom_command/mocks/store.go
	@mockgen -source=./pkg/repo/config_xp_level/store.go -destination=./pkg/repo/config_xp_level/mocks/store.go
	@mockgen -source=./pkg/repo/token/store.go -destination=./pkg/repo/token/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_token/store.go -destination=./pkg/repo/guild_config_token/mocks/store.go
	@mockgen -source=./pkg/repo/discord_guild_stats/store.go -destination=./pkg/repo/discord_guild_stats/mocks/store.go
	@mockgen -source=./pkg/repo/discord_guild_stat_channels/store.go -destination=./pkg/repo/discord_guild_stat_channels/mocks/store.go
	@mockgen -source=./pkg/repo/chain/store.go -destination=./pkg/repo/chain/mocks/store.go
	@mockgen -source=./pkg/service/coingecko/service.go -destination=./pkg/service/coingecko/mocks/service.go
	@mockgen -source=./pkg/repo/nft_sales_tracker/store.go -destination=./pkg/repo/nft_sales_tracker/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_sales_tracker/store.go -destination=./pkg/repo/guild_config_sales_tracker/mocks/store.go
	@mockgen -source=./pkg/service/indexer/service.go -destination=./pkg/service/indexer/mocks/service.go
	@mockgen -source=./pkg/service/marketplace/service.go -destination=./pkg/service/marketplace/mocks/service.go
	@mockgen -source=./pkg/service/abi/service.go -destination=./pkg/service/abi/mocks/service.go
	@mockgen -source=./pkg/repo/nft_collection/store.go -destination=./pkg/repo/nft_collection/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_level_role/store.go -destination=./pkg/repo/guild_config_level_role/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_wallet_verification_message/store.go -destination=./pkg/repo/guild_config_wallet_verification_message/mocks/store.go


setup-githook:
	@echo Setting up softlink pre-commit hooks
	@git config core.hooksPath .githooks/
	@chmod +x .githooks/*
	@echo - done -