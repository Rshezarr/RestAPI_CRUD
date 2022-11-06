build:
	go build -o api_main main.go

dcbuild:
	docker-compose build

up:
	docker-compose up -d

down: 
	docker-compose down