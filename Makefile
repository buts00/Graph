
.PHONY: run
run:
	go run cmd/main.go


.PHONY: build
build:
	go build cmd/main.go


.PHONY: update_packages
update_packages:
	go mod tidy


.PHONY: create_migrations
create_migrations:
	migrate create -dir ./migrations -ext sql -seq init


DB_URL = postgres://postgres:$(POSTGRES_PASSWORD)@localhost:5432/graph_db?sslmode=disable


.PHONY: migrate_up
migrate_up:
	migrate -path ./migrations -database $(DB_URL) up


.PHONY: migrate_down
migrate_down:
	migrate -path ./migrations -database $(DB_URL) down

