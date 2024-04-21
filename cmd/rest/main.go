package main

import (
	db2 "github.com/pangolin-do-golang/tech-challenge/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsn := "host=localhost user=user password=pass dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	orderRepository := db2.NewPostgresOrderRepository(db)
	orderService := order.NewOrderService(orderRepository)

	restServer := server.NewRestServer(orderService)

	restServer.Serve()
}
