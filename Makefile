.PHONY: run build update_packages create_migrations migrate_up migrate_down navigate_python create_venv activate_venv_windows activate_venv_macOs install_requirements run_flask

# Go commands
run:
	go run cmd/main.go

build:
	go build cmd/main.go

update_packages:
	go mod tidy

create_migrations:
	migrate create -dir ./migrations -ext sql -seq init

# Database URL
DB_URL = postgres://postgres:$(POSTGRES_PASSWORD)@localhost:5432/graph_db?sslmode=disable

migrate_up:
	migrate -path ./migrations -database $(DB_URL) up

migrate_down:
	migrate -path ./migrations -database $(DB_URL) down

# Python commands
navigate_python:
	cd graph-recognition

create_venv:
	python3 -m venv venv

activate_venv_windows:
	venv\Scripts\activate

activate_venv_macOs:
	source venv/bin/activate

install_requirements:
	pip install -r requirements.txt

run_flask:
	python3 main.py
