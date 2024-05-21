SHELL := /bin/bash

SERVER_NAME := gopher
SERVER_SOURCE = src
SERVER_APP := $(SERVER_SOURCE)/app.py
BUILD_TAG := latest

.PHONY: dev
dev: build
	sudo docker compose up -d server

.PHONY: build
build: $(SERVER_SOURCE)
	sudo docker build -t $(SERVER_NAME):$(BUILD_TAG) .

.PHONY: down
down:
	sudo docker compose down

.PHONY: admin
admin:
	sudo docker compose up -d admin
