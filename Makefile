# Migration
MIGRATION_NAME ?= $(shell bash -c 'read -p "Migration Name: " m_name; echo $$m_name')
migration_create:
	@clear
	migrate create -ext sql -dir database/migrations $(MIGRATION_NAME)

migration_up_all:
	migrate -database "mysql://root:password@tcp(localhost:3306)/yourdb" -path database/migrations up

migration_up:
	migrate -database "mysql://root:password@tcp(localhost:3306)/yourdb" -path database/migrations up 1

migration_down:
	migrate -database "mysql://root:password@tcp(localhost:3306)/yourdb" -path database/migrations down 1

migration_down_all:
	migrate -database "mysql://root:password@tcp(localhost:3306)/yourdb" -path database/migrations down

migration_reset:
	migrate -database "mysql://root:password@tcp(localhost:3306)/yourdb" -path database/migrations down --all
	migrate -database "mysql://root:password@tcp(localhost:3306)/yourdb" -path database/migrations up
	

# Injector
inject:
	go run github.com/google/wire/cmd/wire

# Lint
lint:
	golangci-lint run

# Main
run:
	go run ./cmd/main.go
build:
	go build main.go


# ==============================================================================
# Docker support

FILES := $(shell docker ps -aq)

down-local:
	docker stop $(FILES)
	docker rm $(FILES)

up-local:
	docker compose up -d

clean:
	docker system prune -f

logs-local:
	docker logs -f $(FILES)