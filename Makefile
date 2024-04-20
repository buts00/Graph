.PHONY: run
run:
	go run cmd/main.go

.PHONY: build
build:
	go build cmd/main.go


.PHONY: download_packages
download_packages:
	go mod tidy