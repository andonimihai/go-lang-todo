# Simple Todo App

This project represents a simple to-do app. It was built by me to learn more about the go-lang. It started with a naive implementation, and I'm adding more features to the project. Current list of the features:

- [Gin Gonic](https://github.com/gin-gonic/gin) Web Framework
- Swagger(with [Gin Swagger](https://github.com/swaggo/gin-swagger)) for API docs
- Postgres as DB with [Gorm](https://gorm.io/index.html) as ORM

## Requirements

- First of all, you need to have [Go](https://go.dev/learn/) installed
- You also need [Docker](https://www.docker.com/get-started) to set up the DB

## Getting started

- Clone the project using the following command `git clone git@github.com:andonimihai/go-lang-todo.git`
- Install all dependencies by running `go mod download.`
- Start the DB by running `docker-compose up -d`(in detach mode)
- Start the project by running `go run main.go`

At this point, the rest API server should update the DB schema and start on http://localhost:3009. You can verify that by running `curl http://localhost:3009/ping`. You should receive the "Pong" message ;)

## Update swagger docs

This project uses [Swag](https://github.com/swaggo/swag) to generate Open API spec(Swagger). You need to install it by using the following command:
`go get -u github.com/swaggo/swag/cmd/swag`
Once installed, you can generate the docs using `swag init --parseDependency --parseInternal` command. For some reason on my WSL it can't find the swag bin. To make it work, I run the following command:
`$HOME/go/bin/swag init --parseDependency --parseInternal`
The swagger docs are available on https://localhost:3009/swagger/index.html
