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
	docker-compose exec app bash

.PHONY: init
init:
	docker-compose exec app go mod init github.com/0kkun/gin-todo-app

.PHONY: server
server:
	docker-compose exec app go run main.go