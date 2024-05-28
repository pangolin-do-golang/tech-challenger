[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=pangolin-do-golang_tech-challenge&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=pangolin-do-golang_tech-challenge) ![Known Vulnerabilities](https://snyk.io/test/github/pcbarretos/pangolin-do-golang/tech-challenge/badge.svg)

# Tech Challenge

## Install

### Go 

- [Go Instalation](https://go.dev/doc/install)

> Make sure you have Go 1.22.2 or higher

Execute 

```shell
go mod tidy
```

## Defining Envs

To correctly use the project, it is necessary to define a .env file, with the values for the envs:

* DB_USERNAME 
* DB_PASSWORD 
* DB_HOST 
* DB_NAME 
* DB_PORT

We recommend using for development the following values:

```
DB_USERNAME=user
DB_PASSWORD=pass
DB_HOST=pgsql
DB_NAME=postgres
DB_PORT=5432
```

## Executing with Docker Compose

```shell
docker-compose build

docker-compose up -d

curl --request GET --url http://localhost:8080/health

## Expected response
= Status Code 200
```

## Accessing Swagger UI

Go to http://localhost:8080/swagger/index.html#/ after the application is running.

## Stack

- [Go](https://go.dev/)
- [Gin Web Framework](https://gin-gonic.com/) - Routes, JSON validation, Error management, Middleware support
- [PostgreSQL](https://www.postgresql.org/) - Database
- [swag](https://github.com/swaggo/swag) - Tool to generate swagger documentation
- [docker](https://www.docker.com/) - Containerization tool
- [docker-compose](https://docs.docker.com/compose/) - Tool to define and run multi-container Docker applications


## DDD with event storm

The team chose to use [Miro](https://miro.com/) to document this deliverable, available at the [link](https://miro.com/app/board/uXjVKVoZwos=/?share_link_id=10494235831).

The diagram contains:

* System documentation in DDD with Event Storm
* Caption for the ubiquitous language used
* Additional details to understand the proposed resolution
* Order fulfillment and payment flow
* Order preparation and delivery flow

## Architecture

This project follows the models proposed by Hexagonal Architecture

![Source: https://making.ndd.tech/reflex%C3%B5es-sobre-o-uso-de-orms-em-dom%C3%ADnios-complexos-parte-2-d7f0ac937121](https://miro.medium.com/v2/resize:fit:521/1*vrXUudR0NzRESXmZl_cHoA.png)

## Swagger

This project makes use of the library [swag](https://github.com/swaggo/swag?tab=readme-ov-file#how-to-use-it-with-gin) to generate the swagger documentation.

### Install

Follow the steps described in the [official documentation](https://github.com/swaggo/swag?tab=readme-ov-file#getting-started)

### Generate 

```shell
 swag init -g cmd/rest/main.go 
```

### Access the documentation

The documentation can be founded at the path `/docs/swagger.yaml` or accessing this [link](./docs/swagger.yaml).

## Infrastructure


[Requierments Infrastructure](https://github.com/pangolin-do-golang/tech-challenge/blob/main/terraform/README.md)

