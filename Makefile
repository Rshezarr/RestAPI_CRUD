build:
	go build -o api_main main.go

dcbuild:
	docker-compose build

up:
	docker-compose up -d

down: 
	docker-compose down

dbrun:
	docker run --name=crud-db -e POSTGRES_PASSWORD='qwerty' -p 5432:5432 -d --rm postgres

ps:
	docker-compose ps