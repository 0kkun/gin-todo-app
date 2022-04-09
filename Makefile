.PHONY: init
init:
	@make build
	@make mod_tidy
	cp .env.example .env

.PHONY: up
up:
	docker-compose up -d

.PHONY: stop
stop:
	docker-compose stop

.PHONY: down
down:
	docker-compose down --remove-orphans

.PHONY: build_c
build_c:
	docker-compose build --no-cache --force-rm

.PHONY: build
build:
	docker-compose build

.PHONY: logs
logs:
	docker-compose logs

.PHONY: app
app:
	docker-compose exec golang bash

.PHONY: mod_init
mod_init:
	go mod init github.com/0kkun/gin-todo-app

.PHONY: mod_tidy
mod_tidy:
	go mod tidy

.PHONY: run
run:
	docker-compose exec golang go run main.go

.PHONY: mysql
mysql:
	docker-compose exec db bash -c 'mysql -u root -p'

.PHONY: db
db:
	docker-compose exec db bash

.PHONY: test
test:
	docker-compose exec golang go test