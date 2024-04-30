# Запускає програму
.PHONY: run
run:
	go run cmd/main.go

# Збирає програму
.PHONY: build
build:
	go build cmd/main.go

# Виконує go mod tidy для оновлення пакетів
.PHONY: update_packages
update_packages:
	go mod tidy

# Створює нову міграцію
.PHONY: create_migrations
create_migrations:
	migrate create -dir ./migrations -ext sql -seq init


DB_URL = postgres://postgres:$(POSTGRES_PASSWORD)@localhost:5432/graph_db?sslmode=disable

# Застосовує міграції до бази даних
.PHONY: migrate_up
migrate_up:
	migrate -path ./migrations -database $(DB_URL) up

# Відміняє міграції у базі даних
.PHONY: migrate_down
migrate_down:
	migrate -path ./migrations -database $(DB_URL) down

