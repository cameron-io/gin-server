SHELL := /bin/bash

SERVER_NAME := gopher
SERVER_APP := main.go
BUILD_TAG := latest

.PHONY: local
local: colima
	$(shell source $(HOME)/.profile)
	go run .

.PHONY: local-db
local-db: colima
	docker compose up -d db

.PHONY: colima
colima:
	colima start

.PHONY: dev
dev: build
	docker compose up -d server

.PHONY: build
build:
	docker build -t $(SERVER_NAME):$(BUILD_TAG) .

.PHONY: down
down:
	docker compose down

.PHONY: clean
clean: down
	docker image rm gopher
	docker volume rm gin_db-data

.PHONY: admin
admin:
	docker compose up -d admin
