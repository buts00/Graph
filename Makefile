.PHONY: run build update_packages create_migrations migrate_up migrate_down navigate_python create_venv activate_venv_windows activate_venv_macOs install_requirements run_flask

# Variables
GO_CMD = go
MIGRATE_CMD = migrate
DB_URL = postgres://postgres:$(POSTGRES_PASSWORD)@localhost:5432/graph_db?sslmode=disable
PYTHON_CMD = python3
VENV_DIR = venv
REQUIREMENTS_FILE = requirements.txt

# Go commands
run:
	$(GO_CMD) run cmd/main.go

build:
	$(GO_CMD) build cmd/main.go

update_packages:
	$(GO_CMD) mod tidy

create_migrations:
	$(MIGRATE_CMD) create -dir ./migrations -ext sql -seq init

migrate_up:
	$(MIGRATE_CMD) -path ./migrations -database $(DB_URL) up

migrate_down:
	$(MIGRATE_CMD) -path ./migrations -database $(DB_URL) down

# Python commands
navigate_python:
	cd graph-recognition

create_venv:
	$(PYTHON_CMD) -m venv $(VENV_DIR)

activate_venv_windows:
	$(VENV_DIR)\Scripts\activate

activate_venv_macOs:
	source $(VENV_DIR)/bin/activate

install_requirements:
	pip install -r $(REQUIREMENTS_FILE)

run_flask:
	$(PYTHON_CMD) main.py
