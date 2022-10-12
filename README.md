# go-rest-api

  This project is my playground for exploring how to use Go as a REST API service by using Gin Web Framework
  
# Stacks

1. [Go](https://go.dev/)  
1. [Gin Web Framework](https://github.com/gin-gonic/gin)
1. [PostgreSQL](https://www.postgresql.org/)
1. [docker-compose](https://docs.docker.com/compose/)
1. [golang-migrate](https://github.com/golang-migrate/migrate)

# Initial setup

1. Install golang-migrate 
    ```
    brew install golang-migrate
    ```
1. Run db
    ```
    docker-compose up -d
    ```
1. Install the UUID OSSP Contrib Module for PostgreSQL
    ```
    docker-compose exec postgres psql -U postgres buddy
    CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
    exit
    ```
1. Migrate db
    ```
    make migrate-up
    ```

# Run Project

1. Run db
    ```
    docker-compose up -d
    ```
1. Run main.go
    ```
    go run ./cmd/main.go
    ```

# Shutdown services

1. Run
    ```
    docker-compose down
    ```

# Add Migration

1. When want to add/alter something in db run
    ```
     make migrate-add 
    ```
