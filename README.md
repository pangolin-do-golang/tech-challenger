# Tech Challenge

## Install

### Go 

- [Go Instalation](https://go.dev/doc/install)

> Make sure you have Go 1.22.2 or higher

Execute 

```shell
go mod tidy
```

### Executing with Docker Compose

```shell
docker-compose build

docker-compose up -d

curl --request GET --url http://localhost:8080/status

## Expected response
{
"status": "ok"
}
```

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