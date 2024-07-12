THIS_FILE := $(lastword $(MAKEFILE_LIST))

.PHONY:  build up start down destroy stop restart logs logs-app ps login-redis login-app db-psql swag test test-cover test-cover-html

help:
	make -pRrq  -f $(THIS_FILE) : 2>/dev/null | awk -v RS= -F: '/^# File/,/^# Finished Make data base/ {if ($$1 !~ "^[#.]") {print $$1}}' | sort | egrep -v -e '^[^[:alnum:]]' -e '^$@$$'

build:
	docker compose  build $(c)

up:
	docker compose  up -d $(c)

start:
	docker compose  start $(c)

down:
	docker compose  down $(c)

destroy:
	docker compose  down -v $(c)

stop:
	docker compose  stop $(c)

restart:
	docker compose  stop $(c)

	docker compose  up -d $(c)

logs:
	docker compose  logs --tail=100 -f $(c)

logs-app:
	docker compose  logs --tail=100 -f app

ps:
	docker compose  ps

login-app:
	docker compose  exec app sh

db-psql:
	docker compose  exec postgres psql -Upostgres

swag: ### initialize or update swag
	swag init -g cmd/app/main.go

test: ### run test
	go test -v ./...

test-cover: ### run test with coverage
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	rm coverage.out

test-cover-html: ### run test with coverage and open html report
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out
	rm coverage.out