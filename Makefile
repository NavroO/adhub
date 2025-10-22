.PHONY: build test run proto lint

build:
	go build ./...

test:
	go test ./...

run:
	go run ./cmd/ads

proto:
	protoc --go_out=proto .

lint:
	golangci-lint run ./...

ci:
	make build
	make test
	make lint

migration-add:
	@if [ -z "$(NAME)" ]; then \
		echo "❌ Please provide migration name using 'make migration-add NAME=create_users_table'"; \
	else \
		migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(NAME); \
	fi

migration-run:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" up

migration-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" down 1

migration-force:
	@read -p "Enter version to force: " version; \
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" force $$version

migration-status:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_URL)" version
