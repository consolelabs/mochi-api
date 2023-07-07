APP_NAME=mochi-api
DEFAULT_PORT=8200
POSTGRES_TEST_CONTAINER?=mochi_local_test

.PHONY: setup init build dev test migrate-up migrate-down

setup:
	go install github.com/rubenv/sql-migrate/...@latest
	go install github.com/golang/mock/mockgen@v1.6.0
	go install github.com/vektra/mockery/v2@latest
	go install github.com/swaggo/swag/cmd/swag@v1.8.6
	go install github.com/cosmtrek/air@latest
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

init-test:
	go install github.com/rubenv/sql-migrate/...@latest
	docker-compose up -d postgres postgres_test redis
	@echo "Waiting for database connection..."
	@while ! docker exec $(POSTGRES_TEST_CONTAINER) pg_isready > /dev/null; do \
		sleep 1; \
	done
	make migrate-test


remove-infras:
	docker-compose down --remove-orphans

build:
	env GOOS=darwin GOARCH=amd64 go build -o bin ./...

dev:
	go run ./cmd/server/main.go

air:
	air -c .air.toml

test:
	make migrate-test
	@PROJECT_PATH=$(shell pwd) go test -cover ./...

migrate-test:
	sql-migrate up -env=test

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
	@mockgen -source=./pkg/repo/discord_guilds/store.go -destination=./pkg/repo/discord_guilds/mocks/store.go
	@mockgen -source=./pkg/repo/config_xp_level/store.go -destination=./pkg/repo/config_xp_level/mocks/store.go
	@mockgen -source=./pkg/repo/token/store.go -destination=./pkg/repo/token/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_token/store.go -destination=./pkg/repo/guild_config_token/mocks/store.go
	@mockgen -source=./pkg/repo/chain/store.go -destination=./pkg/repo/chain/mocks/store.go
	@mockgen -source=./pkg/service/coingecko/service.go -destination=./pkg/service/coingecko/mocks/service.go
	@mockgen -source=./pkg/service/covalent/service.go -destination=./pkg/service/covalent/mocks/service.go
	@mockgen -source=./pkg/repo/guild_config_sales_tracker/store.go -destination=./pkg/repo/guild_config_sales_tracker/mocks/store.go
	@mockgen -source=./pkg/service/indexer/service.go -destination=./pkg/service/indexer/mocks/service.go
	@mockgen -source=./pkg/service/marketplace/service.go -destination=./pkg/service/marketplace/mocks/service.go
	@mockgen -source=./pkg/service/abi/service.go -destination=./pkg/service/abi/mocks/service.go
	@mockgen -source=./pkg/service/discord/service.go -destination=./pkg/service/discord/mocks/service.go
	@mockgen -source=./pkg/service/processor/service.go -destination=./pkg/service/processor/mocks/service.go
	@mockgen -source=./pkg/service/skymavis/service.go -destination=./pkg/service/skymavis/mocks/service.go
	@mockgen -source=./pkg/service/ronin/service.go -destination=./pkg/service/ronin/mocks/service.go
	@mockgen -source=./pkg/repo/nft_collection/store.go -destination=./pkg/repo/nft_collection/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_level_role/store.go -destination=./pkg/repo/guild_config_level_role/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_wallet_verification_message/store.go -destination=./pkg/repo/guild_config_wallet_verification_message/mocks/store.go
	@mockgen -source=./pkg/repo/coingecko_supported_tokens/store.go -destination=./pkg/repo/coingecko_supported_tokens/mocks/store.go
	@mockgen -source=./pkg/repo/user_watchlist_item/store.go -destination=./pkg/repo/user_watchlist_item/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_gm_gn/store.go -destination=./pkg/repo/guild_config_gm_gn/mocks/store.go
	@mockgen -source=./pkg/repo/discord_user_gm_streak/store.go -destination=./pkg/repo/discord_user_gm_streak/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_welcome_channel/store.go -destination=./pkg/repo/guild_config_welcome_channel/mocks/store.go
	@mockgen -source=./pkg/repo/message_repost_history/store.go -destination=./pkg/repo/message_repost_history/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_default_roles/store.go -destination=./pkg/repo/guild_config_default_roles/mocks/store.go
	@mockgen -source=./pkg/repo/guild_config_reaction_roles/store.go -destination=./pkg/repo/guild_config_reaction_roles/mocks/store.go
	@mockgen -source=./pkg/repo/quest/store.go -destination=./pkg/repo/quest/mocks/store.go
	@mockgen -source=./pkg/repo/quest_reward_type/store.go -destination=./pkg/repo/quest_reward_type/mocks/store.go
	@mockgen -source=./pkg/repo/quest_user_log/store.go -destination=./pkg/repo/quest_user_log/mocks/store.go
	@mockgen -source=./pkg/repo/quest_user_list/store.go -destination=./pkg/repo/quest_user_list/mocks/store.go
	@mockgen -source=./pkg/repo/quest_pass/store.go -destination=./pkg/repo/quest_pass/mocks/store.go
	@mockgen -source=./pkg/repo/quest_user_pass/store.go -destination=./pkg/repo/quest_user_pass/mocks/store.go
	@mockgen -source=./pkg/repo/quest_reward/store.go -destination=./pkg/repo/quest_reward/mocks/store.go
	@mockgen -source=./pkg/repo/quest_user_reward/store.go -destination=./pkg/repo/quest_user_reward/mocks/store.go
	@mockgen -source=./pkg/repo/offchain_tip_bot_chain/store.go -destination=./pkg/repo/offchain_tip_bot_chain/mocks/store.go
	@mockgen -source=./pkg/repo/offchain_tip_bot_tokens/store.go -destination=./pkg/repo/offchain_tip_bot_tokens/mocks/store.go
	@mockgen -source=./pkg/repo/offchain_tip_bot_contract/store.go -destination=./pkg/repo/offchain_tip_bot_contract/mocks/store.go
	@mockgen -source=./pkg/repo/offchain_tip_bot_activity_logs/store.go -destination=./pkg/repo/offchain_tip_bot_activity_logs/mocks/store.go
	@mockgen -source=./pkg/repo/offchain_tip_bot_user_balances/store.go -destination=./pkg/repo/offchain_tip_bot_user_balances/mocks/store.go
	@mockgen -source=./pkg/repo/offchain_tip_bot_user_balance_snapshot/store.go -destination=./pkg/repo/offchain_tip_bot_user_balance_snapshot/mocks/store.go
	@mockgen -source=./pkg/repo/offchain_tip_bot_transfer_histories/store.go -destination=./pkg/repo/offchain_tip_bot_transfer_histories/mocks/store.go


setup-githook:
	@echo Setting up softlink pre-commit hooks
	@git config core.hooksPath .githooks/
	@chmod +x .githooks/*
	@echo - done -

gen-swagger:
	swag init  --parseDependency --parseInternal -g ./cmd/server/main.go

pprof:
	rm -f torch.svg && docker run uber/go-torch -u http://host.docker.internal:8200/debug/pprof -p -t=30 > torch.svg