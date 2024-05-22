.PHONY: run build update_packages download_packages create_migrations migrate_up migrate_down navigate_python create_venv activate_venv_windows activate_venv_macOs install_requirements run_flask all_go all_python all

# Go commands
run:
	go run cmd/main.go

build:
	go build cmd/main.go

download_packages:
	go mod download

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

# Combined Go commands
all_go: download_packages build run


create_venv:
	cd graph-recognition && python3 -m venv venv

activate_venv_windows:
	cd graph-recognition && venv\Scripts\activate

activate_venv_macOs:
	cd graph-recognition && source venv/bin/activate

install_requirements:
	cd graph-recognition && venv/bin/pip install -r requirements.txt

run_flask:
	cd graph-recognition && venv/bin/python3 main.py

all_python_mac: create_venv activate_venv_macOs install_requirements run_flask
all_python_windows: create_venv activate_venv_windows install_requirements run_flask