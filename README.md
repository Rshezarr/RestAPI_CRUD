# API

## Description

This application is capable of manipulating information.
In particular, you can: create, read, update and delete a user from the database.

## Usage

### Start Database

1. Open terminal and type command below to start database by docker:

```
$ docker run --name=crud-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 -d --rm postgres

$ docker exec -it crud-db bash
```

2. In purpose to check database fields run command:

```
# psql -U postgres
```

### Start API server

1. Type in terminal in main directory

```
$ go run .
```

Server starts at port 8080. You can create and etc. by using Postman

## Author

Rakhat (@Rshezar)

## Tools used in this project

1. Golang
2. Docker
3. Postgres DB
