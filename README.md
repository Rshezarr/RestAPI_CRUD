# API

## Description

This application is capable of manipulating information.
In particular, you can: create, read, update and delete a user from the database.

## Usage

### Start Database

1. Open terminal and type command below to start database by docker:

```
$ make dbrun
```

2. In purpose to check database fields run command:

```
# docker exec -it $DATABASE_NAME bash

# psql -U postgres
```

### Start API server

1. Type in terminal in main directory

```
$ make run
```

Server starts at port 8080. By using Postman, you can do create, read etc. queries.

## Author

Rakhat (@Rshezarr)

## Tools used in this project

1. Golang
2. Docker
3. Postgres DB
