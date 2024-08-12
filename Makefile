SHELL := /bin/bash

SERVER_NAME := gopher
SERVER_SOURCE = src
SERVER_APP := $(SERVER_SOURCE)/app.py
BUILD_TAG := latest

.PHONY: dev
dev: build
	docker compose up -d server

.PHONY: build
build: $(SERVER_SOURCE)
	docker build -t $(SERVER_NAME):$(BUILD_TAG) .

.PHONY: down
down:
	docker compose down

.PHONY: admin
admin:
	docker compose up -d admin
