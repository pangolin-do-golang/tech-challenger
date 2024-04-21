package main

import (
	db2 "github.com/pangolin-do-golang/tech-challenge/internal/adapters/db"
	"github.com/pangolin-do-golang/tech-challenge/internal/adapters/rest/server"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/order"
	"github.com/pangolin-do-golang/tech-challenge/internal/application/product"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func main() {
	dsn := "host=localhost user=user password=pass dbname=postgres port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Panic(err)
	}

	orderRepository := db2.NewPostgresOrderRepository(db)
	orderService := order.NewOrderService(orderRepository)

	productRepository := db2.NewPostgresProductRepository(db)
	productService := product.NewProductService(productRepository)

	restServer := server.NewRestServer(orderService, productService)

	restServer.Serve()
}
